package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		count := typeCount(s.Text())
		for k, v := range count {
			fmt.Printf("%s  %d\n", k, v)
		}
	}
}

// 使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。
func typeCount(s string) map[string]int {
	var count = map[string]int{}
	for _, v := range s {
		if unicode.IsLetter(v) {
			count["letter"]++
		} else if unicode.IsNumber(v) {
			count["number"]++
		} else if unicode.IsSpace(v) {
			count["space"]++
		} else if unicode.IsPunct(v) {
			count["punct"]++
		} else {
			count["undefined"]++
		}
	}
	return count
}
