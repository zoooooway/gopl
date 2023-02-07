package html

import (
	"fmt"

	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
)

// 编写多参数版本的ElementsByTagName，函数接收一个HTML结点树以及任意数量的标签名，返回与这些标签名匹配的所有元素。
// 简化，改为寻找id
func ElementsByIds(n *html.Node, ids ...string) []*html.Node {
	return forEachNode(n, startElement, endElement, ids)
}

func startElement(n *html.Node, ids []string) bool {
	for _, v := range n.Attr {
		if v.Key == "id" || v.Key == "ID" {
			if slices.Contains(ids, v.Val) {
				return false
			}
		}
	}
	return true
}

func endElement(n *html.Node, ids []string) bool {
	return true
}

func toString(attr []html.Attribute) (comment string) {
	for _, v := range attr {
		comment += fmt.Sprintf("%s=%q", v.Key, v.Val)
	}
	return
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, ids []string) bool, ids []string) (nodes []*html.Node) {
	if pre != nil {
		if !pre(n, ids) {
			nodes = append(nodes, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := forEachNode(c, pre, post, ids)
		if result != nil {
			nodes = append(nodes, result...)
		}
	}
	if post != nil {
		if !post(n, ids) {
			nodes = append(nodes, n)
		}
	}
	return nodes

}
