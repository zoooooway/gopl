package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// 编写函数，记录在HTML树中出现的同名元素的次数。
func main() {
	name := os.Args[1]
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	sum := count(doc, name, 0)
	fmt.Println(sum)
}

func count(node *html.Node, name string, sum int) int {
	if node.Data == name {
		sum += 1
	}
	if node.FirstChild != nil {
		sum = count(node.FirstChild, name, sum)
	}
	if node.NextSibling != nil {
		sum = count(node.NextSibling, name, sum)
	}
	return sum
}
