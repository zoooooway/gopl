package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// todo 给定示例程序是有bug的，不能正确识别自闭合标签
// 如果你选择的id是在自闭合标签上定义的，比如  <h2><a id="c" /></h2>, 那么就找不到这个标签，甚至使用 h2 a这种方式去查找也找不到....
// 妈的调试了半天 ctmd

// 7.17： 扩展xmlselect程序以便让元素不仅可以通过名称选择，也可以通过它们CSS风格的属性进行选择。
// 例如一个像这样
// <div id="page" class="wide">
// 的元素可以通过匹配id或者class，同时还有它的名称来进行选择。
var id = flag.String("i", "", "id args")
var class = flag.String("c", "", "class args")
var ele = elementFlag{}

type elementFlag struct {
	s []string
}

func (e *elementFlag) Set(p string) error {
	e.s = append(e.s, strings.Split(p, " ")...)
	return nil
}
func (e *elementFlag) String() string {
	return fmt.Sprintf("%v", e.s)
}

func fetch(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%d\n", res.StatusCode)
	io.Copy(os.Stdin, res.Body)
	res.Body.Close()
}

func main() {
	flag.CommandLine.Var(&ele, "e", "element args")

	flag.Parse()

	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement // stack of element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if len(stack) > 0 && containsAll(stack, ele.s) {
				buf := bytes.Buffer{}
				buf.WriteString(stack[0].Name.Local)
				for _, v := range stack[1:] {
					buf.WriteString(" ")
					buf.WriteString(v.Name.Local)
				}
				fmt.Printf("%s: %s\n", buf.String(), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []xml.StartElement, y []string) bool {
	f := false
	for _, v := range x {
		if containId(v.Attr) && containClass(v.Attr) {
			f = true
			break
		}
	}
	if !f {
		return f
	}

	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0].Name.Local == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return true
}

func containId(a []xml.Attr) bool {
	if *id == "" {
		return true
	}
	for _, v := range a {
		if v.Name.Local == "id" && v.Value == *id {
			return true
		}
	}

	return false
}

func containClass(c []xml.Attr) bool {
	if *class == "" {
		return true
	}

	for _, v := range c {
		if v.Name.Local == "class" && v.Value == *class {
			return true
		}
	}

	return false
}
