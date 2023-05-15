package main

import (
	"flag"
	"fmt"
	"gopl/ch5/func/eg5.13/links"
	"log"
)

// 练习 8.6： 为并发爬虫增加深度限制。
// 也就是说，如果用户设置了depth=3，那么只有从首页跳转三次以内能够跳到的页面才能被抓取到。
var depth = flag.Uint("depth", 3, "Specify depth")
var tokens = make(chan struct{}, 10)

func main() {

	root := flag.String("root", "", "Crawl root url")
	flag.Parse()

	worklist := make(chan []link)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- []link{{*root, 0}} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, lk := range list {
			if !seen[lk.url] {
				seen[lk.url] = true
				n++
				go func(lk link) {
					lks := crawl(lk)
					worklist <- lks
				}(lk)
			}
		}

	}
	fmt.Println("end")
}

func crawl(l link) []link {
	if l.depth > uint8(*depth) {
		return []link{}
	}
	fmt.Printf("url: %s, depth: %d\n", l.url, l.depth)
	tokens <- struct{}{}
	list, err := links.Extract(l.url)
	<-tokens
	if err != nil {
		log.Print(err)
		return []link{}
	}

	lks := make([]link, len(list))
	d := l.depth + 1
	for i := 0; i < len(list); i++ {
		lks[i] = link{list[i], d}
	}
	return lks
}

type link struct {
	url   string
	depth uint8
}
