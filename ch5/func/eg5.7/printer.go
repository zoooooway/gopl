package html

import (
	"fmt"

	"golang.org/x/net/html"
)

// 完善startElement和endElement函数，使其成为通用的HTML输出器。
// 要求：输出注释结点，文本结点以及每个元素的属性（< a href='...'>）。
// 使用简略格式输出没有孩子结点的元素（即用<img/>代替<img></img>）。
// 编写测试，验证程序输出的格式正确。（详见11章）

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		content := ToString(n.Attr)
		close := ""
		if n.FirstChild == nil {
			close = "/"
		}
		fmt.Printf("%*s<%s %s %s>\n", depth*2, "", n.Data, content, close)
		depth++
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!-- %s", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild == nil {
			return
		}
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	} else if n.Type == html.CommentNode {
		fmt.Print(" -->\n")
	}

}

func ToString(attr []html.Attribute) (comment string) {
	for _, v := range attr {
		comment += fmt.Sprintf("%s=%q", v.Key, v.Val)

	}
	return
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func Print(n *html.Node) {
	forEachNode(n, startElement, endElement)
}
