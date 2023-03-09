package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

/*
8.1： 修改clock2来支持传入参数作为端口号，然后写一个clockwall的程序，
这个程序可以同时与多个clock服务器通信，从多个服务器中读取时间，并且在一个表格中一次显示所有服务器传回的结果，类似于你在某些办公室里看到的时钟墙。
如果你有地理学上分布式的服务器可以用的话，让这些服务器跑在不同的机器上面；
或者在同一台机器上跑多个不同的实例，这些实例监听不同的端口，假装自己在不同的时区。像下面这样：
```
$ TZ=US/Eastern    ./clock2 -port 8010 &
$ TZ=Asia/Tokyo    ./clock2 -port 8020 &
$ TZ=Europe/London ./clock2 -port 8030 &
$ clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
```
*/
type clock struct {
	time string
	zone string
}

func (c clock) String() string {
	return fmt.Sprintf("%s: %s", c.zone, c.time)
}

func main() {
	var clockwall = []*clock{}
	for _, s := range os.Args[1:] {
		l := strings.Split(s, "=")
		ck := &clock{"", l[0]}
		clockwall = append(clockwall, ck)
		conn, err := net.Dial("tcp", l[1])
		if err != nil {
			log.Fatal(err)
		}
		go getTime(conn, ck)
	}
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(clockwall)
	}
}

func getTime(conn net.Conn, c *clock) {
	defer conn.Close()
	scaner := bufio.NewScanner(conn)
	for scaner.Scan() {
		scaner.Text()
		c.time = scaner.Text()
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
