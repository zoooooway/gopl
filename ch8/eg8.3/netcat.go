package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

// 练习 8.3： 在netcat3例子中，conn虽然是一个interface类型的值，但是其底层真实类型是*net.TCPConn，代表一个TCP连接。
// 一个TCP连接有读和写两个部分，可以使用CloseRead和CloseWrite方法分别关闭它们。
// 修改netcat3的主goroutine代码，只关闭网络连接中写的部分，这样的话后台goroutine可以在标准输入被关闭后继续打印从reverb1服务器传回的数据。
func main() {
	addr := flag.String("addr", "localhost:8000", "tcp address")
	flag.Parse()

	raddr, e := net.ResolveTCPAddr("tcp", *addr)
	if e != nil {
		log.Fatal("illegal address")
	}
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	<-done // wait for background goroutine to finish
	conn.Close()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
