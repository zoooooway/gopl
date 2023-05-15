package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 练习 8.9： 编写一个du工具，每隔一段时间将root目录下的目录大小计算并显示出来。
var sema = make(chan struct{}, 50)

func main() {
	log.Println("start...")
	flag.Parse()

	root := flag.Arg(0)

	ch := make(chan int64)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		walkDir(root, ch, &wg)
		wg.Wait()
		close(ch)
	}()
	var filecount int64 = 0
	var bytecount int64 = 0

	tick := time.Tick(1 * time.Second)
loop:
	for {
		select {
		case size, ok := <-ch:
			if !ok {
				log.Printf("%d files, %d M\n", filecount, bytecount/1024/1024)
				break loop
			}
			filecount++
			bytecount += size

		case <-tick:
			log.Printf("%d files, %d M\n", filecount, bytecount/1024/1024)
		}

	}

}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)

			go func() {
				walkDir(subdir, fileSizes, wg)
			}()
		} else {
			fileSizes <- entry.Size()
		}
	}

}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	entries, err := ioutil.ReadDir(dir)
	<-sema
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
