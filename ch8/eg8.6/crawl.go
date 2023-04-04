package main

import (
	"flag"
	"fmt"
	"gopl/ch5/func/eg5.13/links"
	"log"
)

// 练习 8.6： 为并发爬虫增加深度限制。
// 也就是说，如果用户设置了depth=3，那么只有从首页跳转三次以内能够跳到的页面才能被抓取到。
var depth = *flag.Int("depth", 3, "Specify depth")

func main() {

	root := flag.String("root", "", "Crawl root url")
	flag.Parse()

	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- []string{*root} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	var size int
	var list []string
	for ; n > 0 && depth > 0; n-- {
		list = append(list, <-worklist...)
		size += len(list)
		if n == 1 {
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string) {
						l := crawl(link)
						if len(l) == 0 {
							return
						}
						worklist <- l
					}(link)
				}
			}
			list = []string{}
			size = 0
			depth--
		}
	}
	fmt.Println("end")
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
