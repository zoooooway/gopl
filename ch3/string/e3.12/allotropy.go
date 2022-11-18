package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {

	fmt.Println(isAllotropy(os.Args[1], os.Args[2]))
}

func isAllotropy(str1 string, str2 string) bool {
	if utf8.RuneCountInString(str1) != utf8.RuneCountInString(str2) {
		return false
	}

	m := make(map[rune]int)
	for _, v := range str1 {
		m[v]++
	}

	for _, v := range str2 {
		if m[v] <= 0 {
			return false
		}
		m[v]--
	}
	return true
}
