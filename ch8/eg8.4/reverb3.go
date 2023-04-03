package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

// 练习8.4： 修改reverb2服务器，在每一个连接中使用sync.WaitGroup来计数活跃的echo goroutine。
// 当计数减为零时，关闭TCP连接的写入，像练习8.3中一样。
// 验证一下你的修改版netcat3客户端会一直等待所有的并发“喊叫”完成，即使是在标准输入流已经关闭的情况下。
func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	f := 0
	for input.Scan() {
		wg.Add(1)
		if f == 0 {
			f = 1
			go func() {
				wg.Wait()
				tc := c.(*net.TCPConn)
				tc.CloseWrite()
				log.Println("CloseWrite...")
			}()
		}
		go echo(c, input.Text(), 1*time.Second, &wg)
	}
	// NOTE: ignoring potential errors from input.Err()

	c.Close()

}
