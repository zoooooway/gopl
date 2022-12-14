package main

import "fmt"

func main() {
	x := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	ret := rotate(x[:], 2, true)
	fmt.Println(ret)
}

// 编写一个rotate函数，通过一次循环完成旋转。
// len为前len个需要旋转的元素 right为是否向右旋转
func rotate(x []int, len int, right bool) []int {
	r := x[len:]
	if right {
		for i := 0; i < len; i++ {
			r = append(r, x[i])
		}
	} else {
		for i := len - 1; i >= 0; i-- {
			r = append(r, x[i])
		}
	}

	return r
}
