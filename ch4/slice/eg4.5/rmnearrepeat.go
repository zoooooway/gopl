package main

import "fmt"

func main() {
	x := [...]string{"h", "e", "l", "l", "o", "w", "o", "r", "l", "d"}
	fmt.Println(rmNearRepeat(x[:]))
}

// 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
func rmNearRepeat(x []string) []string {
	if len(x) == 1 {
		return x
	}
	c := 1
	i := 0
	for c < len(x) {
		if x[i] == x[c] {
			c++
			continue
		}
		x[i+1] = x[c]
		i++
		c++
	}
	return x[:i+1]
}
