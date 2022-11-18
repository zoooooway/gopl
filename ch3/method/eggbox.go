package method

import "math"

// 我不理解这个生成图形函数的原理，抄的: https://github.com/PieerePi/gople/blob/master/ch3/e3.2/surface.go
func Eggbox(x, y float64) float64 {
	return -0.1 * (math.Cos(x) + math.Cos(y))
}
