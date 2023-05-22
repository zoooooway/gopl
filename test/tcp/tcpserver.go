package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

// 简单的tcp服务端，只会打印收到的数据
func main() {
	port := flag.String("port", "8010", "tcp server listen port")
	ip := flag.String("ip", "", "tcp server listen ip")
	flag.Parse()

	l, err := net.Listen("tcp", *ip+":"+*port)
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

func handleConn(c net.Conn) {
	client := c.RemoteAddr().Network() + ":" + c.RemoteAddr().String()
	input := bufio.NewScanner(c)
	for input.Scan() {
		// only print message
		fmt.Printf("%s> %s\n", client, input.Text())
	}
	// NOTE: ignoring potential errors from input.Err()

	c.Close()
}
