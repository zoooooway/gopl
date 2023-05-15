package main

import (
	"flag"
	"fmt"
	"gopl/ch5/func/eg5.13/links"
	"log"
	"os"
)

// 练习 8.10： HTTP请求可能会因http.Request结构体中Cancel channel的关闭而取消。
// 修改8.6节中的web crawler来支持取消http请求。
// 提示：http.Get并没有提供方便地定制一个请求的方法。
// 你可以用http.NewRequest来取而代之，设置它的Cancel字段，然后用http.DefaultClient.Do(req)来进行这个http请求。
var depth = flag.Uint("depth", 3, "Specify depth")
var tokens = make(chan struct{}, 10)
var done = make(chan struct{})

func main() {

	root := flag.String("root", "", "Crawl root url")
	flag.Parse()

	// 监听关闭输入
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		log.Println("cancel.....")
		close(done)
	}()

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

	if canceled() {
		return []link{}
	}
	fmt.Printf("url: %s, depth: %d\n", l.url, l.depth)
	tokens <- struct{}{}
	list, err := links.Extract2(l.url, done)
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

func canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
