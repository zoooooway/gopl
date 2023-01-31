package count

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// 实现countWordsAndImages。（参考练习4.9如何分词）
func countWordsAndImages(url string) (wc, ic int) {
	r := fetch(url)
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	// count words
	text := textCollect(doc)
	fmt.Println(text)
	for _, v := range text {
		words := strings.Split(v, " ")
		wc += len(words)
	}

	// count images
	imgs := imgCollect(doc)
	ic += len(imgs)

	return
}

func fetch(url string) io.ReadCloser {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%d\n", res.StatusCode)
	return res.Body
}

func textCollect(node *html.Node) (text []string) {
	if node.Type == html.TextNode && strings.TrimSpace(node.Data) != "" {
		if node.Parent != nil && node.Parent.Type == html.ElementNode && node.Parent.Data != "script" && node.Parent.Data != "style" {
			text = append(text, node.Data)
		}
	}

	if node.FirstChild != nil {
		text = append(text, textCollect(node.FirstChild)...)
	}
	if node.NextSibling != nil {
		text = append(text, textCollect(node.NextSibling)...)
	}
	return
}

func imgCollect(node *html.Node) (imgs []string) {
	if node.Type == html.ElementNode && strings.TrimSpace(node.Data) == "img" {
		imgs = append(imgs, node.Data)
	}

	if node.FirstChild != nil {
		imgs = append(imgs, textCollect(node.FirstChild)...)
	}
	if node.NextSibling != nil {
		imgs = append(imgs, textCollect(node.NextSibling)...)
	}
	return
}
