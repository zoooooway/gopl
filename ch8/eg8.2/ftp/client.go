package ftp

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

type FtpRequest struct {
	cmd string
	f   []byte
}

var conn net.Conn

func DefaultClient() {
	http.HandleFunc("/ftp", handleCMD)
	http.HandleFunc("/ftp/connect", handleConnect)

	http.ListenAndServe("localhost:8001", nil)

	log.Panicln("sever stop")

	if conn != nil {
		conn.Close()
	}
}

func handleConnect(w http.ResponseWriter, r *http.Request) {
	c, e := net.Dial("tcp", "127.0.0.1:8022")
	if e != nil {
		io.WriteString(w, e.Error())
		return
	}
	conn = c
	w.Write([]byte("connected to 127.0.0.1:8022"))
}

func handleCMD(w http.ResponseWriter, r *http.Request) {
	pms := r.URL.Query()

	switch pms.Get("cmd") {
	case "cd":
		dir := pms.Get("dir")

		io.WriteString(conn, fmt.Sprintf("cd %s\n", dir))
		bts, e := io.ReadAll(conn)
		if e != nil {
			w.Write([]byte(e.Error()))
		}
		io.WriteString(w, string(bts))
	case "ls":
		io.WriteString(conn, "ls\n")
		log.Println("尝试读取响应")
		bts, e := io.ReadAll(conn)
		log.Println("读取响应结束")
		if e != nil {
			w.Write([]byte(e.Error()))
		}
		io.WriteString(w, string(bts))
	case "pwd":
		io.WriteString(conn, "pwd\n")
		bts, e := io.ReadAll(conn)
		if e != nil {
			w.Write([]byte(e.Error()))
		}
		io.WriteString(w, string(bts))
	case "get":
		file := pms.Get("file")

		i := strings.LastIndex(file, string(os.PathSeparator))
		if i == -1 {
			io.WriteString(w, "illegal input")
			return
		}

		io.WriteString(conn, fmt.Sprintf("get %s\n", file))

		fname := file[i:]
		f, e := os.CreateTemp(".", "*"+fname)
		if e != nil {
			w.Write([]byte(e.Error()))
			return
		}

		var pos int64
		bts := make([]byte, 1024)
		for {
			log.Println("尝试读取响应...")
			n, e := conn.Read(bts)
			if e != nil {
				if e == io.EOF {
					break
				}
				w.Write([]byte(e.Error()))
				return
			}
			log.Printf("尝试写入文件...%d\n", n)
			f.WriteAt(bts, pos)
			pos += int64(n)
		}
		log.Println("读取完毕")
		b, e := io.ReadAll(f)
		if e != nil {
			w.Write([]byte(e.Error()))
			return
		}
		w.Write(b)

	case "send":
	case "close":
		log.Printf("conn: %s close!\n", conn.RemoteAddr().String())
		return
	}
}
