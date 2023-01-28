package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// 修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit1(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit1(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit1(links, c)
	}
	return links
}

// recursion version
func visit2(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.FirstChild != nil {
		links = visit2(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit2(links, n.NextSibling)
	}

	return links
}
