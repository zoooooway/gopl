package crawl

import (
	"log"
	"testing"
)

func TestCrawl(t *testing.T) {
	urls, e := crawl2("https://golang.google.cn/")
	if e != nil {
		log.Println(e)
	}
	log.Println(urls)
}
