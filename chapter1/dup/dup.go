package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// input.Scan()：读入输入流的下一行，内容可以通过input.Test()获取到，到读取不到内容时返回false
	// 读取标准输入流时按ctrl+z 加回车可以终止输入
	for input.Scan() {
		line := input.Text()
		fmt.Println(line)
		count[line]++
	}
	for line, n := range count {
		if n > 1 {
			fmt.Printf("%d, %s\n", n, line)
		}

	}

}
