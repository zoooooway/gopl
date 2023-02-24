package parser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值（和其它值）。
// 实现一个简单版本的NewReader，用它来构造一个接收字符串输入的HTML解析器（§5.2）

type Parser interface {
	io.Reader
	Parse() (*html.Node, error)
}

type HtmlParser struct {
	io.Reader
}

func (hp HtmlParser) Read(p []byte) (n int, err error) {
	return hp.Reader.Read(p)
}

func (hp HtmlParser) Parse() (*html.Node, error) {
	return html.Parse(hp.Reader)
}

func NewReader(s string) Parser {
	r := strings.NewReader(s)
	hp := HtmlParser{r}
	return hp
}
