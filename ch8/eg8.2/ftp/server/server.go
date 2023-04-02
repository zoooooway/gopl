package server

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
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

	for {
		str, e := bufio.NewReader(c).ReadString('\n')
		if e != nil {
			io.WriteString(c, e.Error())
			io.WriteString(c, "\n")
			continue
		}

		// log.Println("exec " + str)
		input := strings.Split(strings.TrimSpace(str), " ")

		cmd := input[0]
		switch cmd {
		case "cd":
			f, e := os.Open(input[1])
			defer f.Close()
			if e != nil {
				io.WriteString(c, e.Error())
				io.WriteString(c, "\n")
				continue
			}

			currPath = input[1]
			io.WriteString(c, currPath)
			io.WriteString(c, "\n")
		case "ls":
			f, e := os.Open(currPath)
			defer f.Close()
			if e != nil {
				io.WriteString(c, e.Error())
				io.WriteString(c, "\n")
				continue
			}

			fs, e := f.Readdir(0)
			if e != nil {
				io.WriteString(c, e.Error())
				io.WriteString(c, "\n")
				continue
			}

			files := []string{}
			for _, v := range fs {
				if v.IsDir() {
					files = append(files, "/"+v.Name())
				} else {
					files = append(files, v.Name())
				}
			}

			ls := strings.Join(files, "\t")
			io.WriteString(c, ls)
			io.WriteString(c, "\n")
		case "pwd":
			io.WriteString(c, currPath)
			io.WriteString(c, "\n")
		case "get":
			fname := input[1]
			download(fname, c)
		case "send":
			fname := input[1]
			update(fname, c)
		case "close":
			log.Printf("conn: %s close!\n", c.RemoteAddr())
			return
		default:
			io.WriteString(c, "UNKNOWN CMD\n")
		}
	}
}

func download(filename string, c net.Conn) {
	f, e := os.OpenFile(filename, os.O_RDONLY, os.ModeExclusive)
	if e != nil {
		io.WriteString(c, e.Error())
		io.WriteString(c, "\n")
		return
	}
	defer f.Close()

	fs, e := f.Stat()
	if e != nil {
		io.WriteString(c, e.Error())
		io.WriteString(c, "\n")
		return
	}
	total := fs.Size()
	c.SetWriteDeadline(time.Now().Add(6 * time.Second))
	c.Write([]byte(strconv.FormatInt(total, 10)))
	io.WriteString(c, "\n")

	w := bufio.NewWriter(c)
	n, e := w.ReadFrom(f)

	if e != nil {
		io.WriteString(c, e.Error())
		io.WriteString(c, "\n")
		return
	}
	log.Printf("total size: %d\n", n)
}

func update(filename string, c net.Conn) {
	name := strings.Split(filename, ".")

	f, e := os.CreateTemp(".", name[0]+"*."+name[1])
	if e != nil {
		log.Println(e.Error())
		io.WriteString(c, e.Error())
		io.WriteString(c, "\n")
		return
	}

	defer f.Close()

	r := bufio.NewReader(c)
	str, e := r.ReadString('\n')
	if e != nil {
		io.WriteString(c, e.Error())
		io.WriteString(c, "\n")
		return
	}

	total, e := strconv.Atoi(str[:len(str)-1])
	log.Printf("total: %d\n", total)
	if e != nil {
		io.WriteString(c, e.Error())
		io.WriteString(c, "\n")
		return
	}

	var counter int64
	bts := make([]byte, 1024)

	for counter != int64(total) {
		c.SetWriteDeadline(time.Now().Add(5 * time.Second))
		n, e := r.Read(bts)
		log.Printf("读取: %d\n", n)
		if e != nil {
			io.WriteString(c, e.Error())
			return
		}

		f.WriteAt(bts, counter)
		counter += int64(n)
	}

	log.Printf("total size: %d\n", counter)
}
