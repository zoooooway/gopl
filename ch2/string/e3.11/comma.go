package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("rest : %s", comma("-1234.23456789"))
}

// 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。
func comma(str string) string {
	var buf bytes.Buffer
	// 浮点数的逗号只会加在整数部
	n := strings.LastIndex(str, ".")

	// 计算第一个逗号的位置
	f := (n - 1) % 3

	if f == 0 {
		f += 3
	}

	for i, v := range str {
		// 如果首位为正负号，那么逗号位置向后移动一位
		if i == 1 && (strings.HasPrefix(str, "-") || strings.HasPrefix(str, "+")) {
			writeRune(buf, v)
			f++
			continue
		}

		if i == f && i < n {
			buf.WriteString(",")
			f += 3
		}
		writeRune(buf, v)
	}
	return buf.String()
}

// write rune to buf.
// program terminates immediately if any error occurs.
func writeRune(buf bytes.Buffer, val rune) {
	_, err := buf.WriteRune(val)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %s", err)
		os.Exit(1)
	}
}
