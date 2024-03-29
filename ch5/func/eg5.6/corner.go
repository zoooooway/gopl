package corner

import "math"

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// 修改gopl.io/ch3/surface (§3.2) 中的corner函数，将返回值命名，并使用bare return。
func corner(i, j int) (sx, sy, z float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z = f(x, y)
	// z := method.Eggbox(x, y)

	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
