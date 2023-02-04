package crawl

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"gopl/ch5/func/eg5.13/links"
)

// 修改crawl，使其能保存发现的页面，必要时，可以创建目录来保存这些页面。
// 只保存来自原始域名下的页面。假设初始页面在golang.org下，就不要保存vimeo.com下的页面。
func crawl2(url string) ([]string, error) {
	fmt.Println(url)
	domain, e := getDomain(url)
	if e != nil {
		return nil, e
	}
	fmt.Printf("domain: %s\n", domain)

	list, err := links.Extract(url)
	if err != nil {
		return nil, err
	}
	// save
	e = os.Mkdir("./pages", fs.ModeDir)
	if e != nil {
		log.Print(e)
	}

	f, e := os.Create("./pages/" + domain + ".txt")
	if e != nil {
		log.Print(e)
	}

	var saveUrl []string
	for _, v := range list {
		d, e := getDomain(v)
		if e != nil {
			continue
		}
		if d == domain {
			saveUrl = append(saveUrl, v)
		}
	}

	pages := strings.Join(saveUrl, "\n")
	f.Write([]byte(pages))

	return saveUrl, nil
}

func getDomain(url string) (string, error) {
	begin := strings.Index(url, "//")
	if begin == -1 {
		return "", errors.New("illegal url")
	}
	sub := url[begin+2:]
	end := strings.Index(sub, "/")
	if end != -1 {
		sub = sub[0:end]
	}
	return sub, nil
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
