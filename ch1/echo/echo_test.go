package main

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkStringAppend(b *testing.B) {
	arg := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	s, sep := "", "----"
	fmt.Println(s)

	for i := 0; i < b.N; i++ {
		// for _, v := range arg {
		// 	s += v + sep
		// }
		// fmt.Println(s)
		// s = ""

		fmt.Println(strings.Join(arg[0:], sep))

	}

}
