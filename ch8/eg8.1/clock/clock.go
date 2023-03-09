package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	p := flag.Int("port", 8000, "assign port")
	tz := flag.String("timezone", "UTC", "assign timezone")

	flag.Parse()
	addr := "localhost:" + fmt.Sprint(*p)
	fmt.Println(addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, *tz) // handle one connection at a time
	}
}

func handleConn(c net.Conn, tz string) {
	fmt.Printf("conn: %s accept!\n", c.RemoteAddr())
	defer c.Close()
	loc, e := time.LoadLocation(tz)
	if e != nil {
		log.Fatalf("%s is not a valid timezone", tz)
	}
	for {
		t := time.Now()
		nt := t.In(loc)
		s := nt.Format("15:04:05\n")
		_, err := io.WriteString(c, s)
		fmt.Printf("write: %s\n", s)
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
