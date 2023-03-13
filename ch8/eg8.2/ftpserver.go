package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// 练习 8.2： 实现一个并发FTP服务器。
// 服务器应该解析客户端发来的一些命令，比如cd命令来切换目录，ls来列出目录内文件，get和send来传输文件，close来关闭连接。
// 你可以用标准的ftp命令来作为客户端，或者也可以自己实现一个。
func main() {
	listener, e := net.Listen("tcp", "127.0.0.1:8022")
	if e != nil {
		log.Fatal("tcp server start failed!", e.Error())
	}

	for {
		conn, e := listener.Accept()
		if e != nil {
			log.Print(e.Error())
			continue
		}
		go handleFtp(conn)
	}
}

func handleFtp(c net.Conn) {
	log.Printf("conn: %s accept!\n", c.RemoteAddr())
	defer c.Close()

	currPath := "/"
	scaner := bufio.NewScanner(c)
	for scaner.Scan() {
		in := scaner.Text()
		input := strings.Split(in, " ")

		cmd := input[0]
		switch cmd {
		case "cd":
			currPath = input[1]
		case "ls":
			f, e := os.Open(currPath)
			if e != nil {
				io.WriteString(c, e.Error())
				continue
			}

		case "pwd":
			io.WriteString(c, currPath)
		case "get":
		case "send":
		case "close":
			log.Printf("conn: %s close!\n", c.RemoteAddr())
			return
		}
	}
}
