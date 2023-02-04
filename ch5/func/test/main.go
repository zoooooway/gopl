package main

import (
	"fmt"
)

// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {
	f1 := squares()
	f2 := squares()
	fmt.Println(f1()) // "1"
	fmt.Println(f1()) // "4"
	fmt.Println(f1()) // "9"
	fmt.Println(f1()) // "16"

	fmt.Println("---------------------")

	fmt.Println(f2()) // "1"
	fmt.Println(f2()) // "4"
	fmt.Println(f2()) // "9"
	fmt.Println(f2()) // "16"
}
