package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	// fmt.Println(string(rmNearSpace([]byte("hello i am lily!  how are you?"))))
	fmt.Println(string(rmNearSpace([]byte("hello i am 李华!  你 咋  样?"))))
}

// 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
func rmNearSpace(x []byte) []byte {
	r := bytes.Runes(x)

	if len(r) == 1 {
		return x
	}
	c := 1
	i := 0
	for c < len(r) {
		fmt.Printf("r[%d]: %c  r[%d]: %c\n", i, r[i], c, r[c])
		if unicode.IsSpace(r[i]) {
			if r[i] == r[c] {
				c++
				continue
			}
		}
		r[i+1] = r[c]
		i++
		c++
	}
	rs := r[:i+1]
	return []byte(string(rs))
}
