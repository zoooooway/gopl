package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// 练习8.8： 使用select来改造8.3节中的echo服务器，为其增加超时，这样服务器可以在客户端10秒中没有任何喊话时自动断开连接。
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
		log.Printf("%s accepted", conn.RemoteAddr())
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
	ticker := time.NewTicker(10 * time.Second)
	alive := &flag{false}
	filepath.Join()
	go func() {
		for {
			select {
			case <-ticker.C:
				if !alive.bool {
					c.Close()
					ticker.Stop()
					return
				}
				alive.bool = false
			}
		}
	}()

	doHandle(c, alive)
}

func doHandle(c net.Conn, alive *flag) {
	defer c.Close()

	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	wg.Add(1)
	for input.Scan() {
		wg.Add(1)
		go func() {
			wg.Wait()
			tc := c.(*net.TCPConn)
			tc.CloseWrite()
			log.Println("CloseWrite...")
		}()
		go func() {
			alive.bool = true
			echo(c, input.Text(), 1*time.Second, &wg)
		}()
	}
	log.Println("close connection..")
	// NOTE: ignoring potential errors from input.Err()
}

type flag struct {
	bool
}
