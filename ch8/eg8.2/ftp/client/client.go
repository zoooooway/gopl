package client

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

	log.Println("sever stop")

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
		res, e := bufio.NewReader(conn).ReadString('\n')
		if e != nil {
			io.WriteString(w, e.Error())
			return
		}
		io.WriteString(w, res)
	case "ls":
		io.WriteString(conn, "ls\n")
		res, e := bufio.NewReader(conn).ReadString('\n')
		if e != nil {
			io.WriteString(w, e.Error())
			return
		}
		io.WriteString(w, res)
	case "pwd":
		io.WriteString(conn, "pwd\n")
		res, e := bufio.NewReader(conn).ReadString('\n')
		if e != nil {
			io.WriteString(w, e.Error())
			return
		}
		io.WriteString(w, res)
	case "get":
		file := pms.Get("file")
		download(file, conn, w)
	case "send":
		filename := pms.Get("file")
		update(filename, r, conn)
	case "close":
		log.Printf("conn: %s close!\n", conn.RemoteAddr().String())
		return
	}
	log.Println("done")
}

func download(file string, conn net.Conn, w http.ResponseWriter) {
	i := strings.LastIndex(file, string(os.PathSeparator))
	if i == -1 {
		io.WriteString(w, "illegal input")
		return
	}

	io.WriteString(conn, fmt.Sprintf("get %s\n", file))

	fname := file[i+1:]
	log.Println(fname)
	f, e := os.CreateTemp(".", "*"+fname)
	if e != nil {
		io.WriteString(w, e.Error())
		return
	}
	defer f.Close()

	r := bufio.NewReader(conn)
	s, e := r.ReadString('\n')

	if e != nil {
		io.WriteString(w, e.Error())
		return
	}

	total, e := strconv.ParseInt(s[:len(s)-1], 10, 0)

	if e != nil {
		io.WriteString(w, e.Error())
		return
	}

	var counter int64
	bts := make([]byte, 1024)

	for counter != total {
		conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
		n, e := r.Read(bts)
		if e != nil {
			io.WriteString(w, e.Error())
			return
		}

		f.WriteAt(bts, counter)
		counter += int64(n)
	}
	log.Println("读取完毕")
}

func update(filename string, r *http.Request, conn net.Conn) {
	io.WriteString(conn, fmt.Sprintf("send %s\n", filename))

	body := r.Body
	defer body.Close()

	w := bufio.NewWriter(conn)
	bts, e := io.ReadAll(body)
	if e != nil {
		io.WriteString(w, e.Error())
		return
	}

	io.WriteString(w, strconv.Itoa(len(bts)))
	io.WriteString(w, "\n")

	nn, e := w.Write(bts)
	if e != nil {
		io.WriteString(w, e.Error())
		return
	}
	w.Flush()
	log.Printf("上传文件大小：%d\n", nn)
}
