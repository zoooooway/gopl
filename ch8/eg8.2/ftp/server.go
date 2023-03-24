package ftp

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// 练习 8.2： 实现一个并发FTP服务器。
// 服务器应该解析客户端发来的一些命令，比如cd命令来切换目录，ls来列出目录内文件，get和send来传输文件，close来关闭连接。
// 你可以用标准的ftp命令来作为客户端，或者也可以自己实现一个。
func DefaultServer() {
	port := flag.String("port", "8022", "ftp server port")
	flag.Parse()

	listener, e := net.Listen("tcp", "127.0.0.1:"+*port)
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

	currPath := "H:\\hzw\\go\\gopl\\ch8\\eg8.2"

	scaner := bufio.NewScanner(c)
	for scaner.Scan() {
		in := scaner.Text()
		log.Println("exec" + in)
		input := strings.Split(in, " ")

		cmd := input[0]
		switch cmd {
		case "cd":
			f, e := os.Open(input[1])
			defer f.Close()
			if e != nil {
				io.WriteString(conn, e.Error())
				continue
			}

			currPath = input[1]
			io.WriteString(conn, currPath)
		case "ls":
			f, e := os.Open(currPath)
			defer f.Close()
			if e != nil {
				io.WriteString(c, e.Error())
				continue
			}

			fs, e := f.Readdir(0)
			if e != nil {
				io.WriteString(c, e.Error())
				continue
			}

			files := []string{}
			for _, v := range fs {
				files = append(files, v.Name())
			}

			ls := strings.Join(files, "\t")
			io.WriteString(c, ls)
			io.WriteString()
			log.Println("发送完毕")
		case "pwd":
			io.WriteString(c, currPath)
		case "get":
			fname := input[1]
			f, e := os.CreateTemp(".", fname)
			defer f.Close()
			if e != nil {
				io.WriteString(c, e.Error())
				continue
			}

			c.LocalAddr()

			r := io.TeeReader(f, c)
			bt := []byte{}

			if _, e := r.Read(bt); e != nil {
				io.WriteString(c, e.Error())
				continue
			}
		case "send":
			fname := input[1]
			f, e := os.OpenFile(fname, os.O_RDONLY, os.ModeExclusive)
			defer f.Close()
			if e != nil {
				io.WriteString(c, e.Error())
				continue
			}

			r := io.TeeReader(f, c)
			bt := []byte{}
			if _, e := r.Read(bt); e != nil {
				io.WriteString(c, e.Error())
				continue
			}

		case "close":
			log.Printf("conn: %s close!\n", c.RemoteAddr())
			return
		}
	}
	log.Printf("conn: %s close!\n", c.RemoteAddr())
}
