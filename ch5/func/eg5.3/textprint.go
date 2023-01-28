package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// 编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素,因为这些元素对浏览者是不可见的。
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	print(doc)
}

func print(node *html.Node) {
	if node.Type == html.TextNode && strings.TrimSpace(node.Data) != "" {
		if node.Parent != nil && node.Parent.Type == html.ElementNode && node.Parent.Data != "script" && node.Parent.Data != "style" {
			fmt.Println(node.Data)
		}

	}

	if node.FirstChild != nil {
		print(node.FirstChild)
	}
	if node.NextSibling != nil {
		print(node.NextSibling)
	}
}
