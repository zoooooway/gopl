package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	x := []byte("哟！这不是jack嘛!")
	reverseForBytes(x)
	fmt.Println(string(x))
}

// 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？
func reverseForBytes(x []byte) {
	if len(x) < 1 {
		return
	}
	r, size := utf8.DecodeRune(x)
	fmt.Printf("decode rune is : %c, size: %d\n", r, size)

	if r == utf8.RuneError {
		panic("Rune Error")
	}

	rotate(x, size)
	reverseForBytes(x[0 : len(x)-size])

}

func rotate(x []byte, len int) {
	reverse(x[:len])
	reverse(x[len:])
	reverse(x)
}

func reverse(x []byte) {
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
}
