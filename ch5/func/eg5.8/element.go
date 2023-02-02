package html

import (
	"fmt"

	"golang.org/x/net/html"
)

// 修改pre和post函数，使其返回布尔类型的返回值。
// 返回false时，中止forEachNode的遍历。
// 使用修改后的代码编写ElementByID函数，根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历。

var depth int

func ElementByID(n *html.Node, id string) *html.Node {
	return forEachNode(n, id, startElement, endElement)

}

func startElement(n *html.Node, id string) bool {
	for _, v := range n.Attr {
		if v.Key == "id" || v.Key == "ID" {
			if v.Val == id {
				return false
			}
		}
	}
	return true
}

func endElement(n *html.Node, id string) bool {
	for _, v := range n.Attr {
		if v.Key == "id" || v.Key == "ID" {
			if v.Val == id {
				return false
			}
		}
	}
	return true
}

func toString(attr []html.Attribute) (comment string) {
	for _, v := range attr {
		comment += fmt.Sprintf("%s=%q", v.Key, v.Val)

	}
	return
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		if !pre(n, id) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := forEachNode(c, id, pre, post)
		if result != nil {
			return result
		}
	}
	if post != nil {
		if !post(n, id) {
			return n
		}
	}
	return nil

}
