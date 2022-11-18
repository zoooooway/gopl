package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("rest : %s\n", comma("1234"))
	fmt.Printf("rest : %s\n", comma("你好啊!this世界"))
}

// 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。
// 函数的功能是将一个表示整数值的字符串，每隔三个字符插入一个逗号分隔符，例如“12345”处理后成为“12,345”。
// 该函数支持汉字
func comma(str string) string {
	var buf bytes.Buffer
	arr := []rune(str)
	// 计算第一个逗号分隔符的位置
	l := len(arr)
	f := l % 3
	if f == 0 {
		f += 3
	}
	c := 0
	// 每次循环读取v值时，索引的移动步长是该字符的字节数，索引不等于遍历数，因此需要一个额外的变量来确认遍历次数
	for _, v := range str {
		if c == f {
			buf.WriteString(",")
			f += 3
		}
		_, err := buf.WriteRune(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %s", err)
			os.Exit(1)
		}
		c++
	}
	return buf.String()
}
