package visit

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// 扩展visit函数，使其能够处理其他类型的结点，如images、scripts和style sheets。
func Visit(contents []string, nType string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == nType {
		for _, a := range n.Attr {
			if a.Key == "href" {
				contents = append(contents, a.Val)
			}
		}
	} else if n.Parent != nil && n.Parent.Data == nType {
		contents = append(contents, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		contents = Visit(contents, nType, c)
	}
	return contents
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
