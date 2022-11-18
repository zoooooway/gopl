// 参考https://github.com/PieerePi/gople/blob/master/ch3/e3.3/surface.go
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/surface", func(w http.ResponseWriter, r *http.Request) {
		// 需设置对应文件的响应体
		w.Header().Set("Content-Type", "image/svg+xml")

		fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)

		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay, az := corner(i+1, j)
				bx, by, bz := corner(i, j)
				cx, cy, cz := corner(i, j+1)
				dx, dy, dz := corner(i+1, j+1)

				// 跳过无效的多边形
				if math.IsInf(az+bz+cz+dz, 0) {
					continue
				}
				if az < 0 && bz < 0 && cz < 0 && dz < 0 {
					fmt.Fprintf(w, "<polygon style='fill: #0000ff;' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
						ax, ay, bx, by, cx, cy, dx, dy)
				} else {
					fmt.Fprintf(w, "<polygon style='fill: #ff0000;' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
						ax, ay, bx, by, cx, cy, dx, dy)
				}

			}
		}
		fmt.Fprintf(w, "</svg>")
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	// z := method.Eggbox(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z

}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
