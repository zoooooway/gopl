package main

import "fmt"

func main() {
	str := "你好"
	fmt.Println(len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("%v\n", str[i])
	}
	// 自动隐式解码UTF8字符串
	// 对于非ASCII，索引更新的步长将超过1个字节
	for i, c := range str {
		fmt.Printf(" %d %c\n", i, c)
	}

	conversation := `
	hello!
	hi!
	nice to meet you
	nice to meet you too
	
			[2022-11-11]
	`
	fmt.Println(conversation)
}
