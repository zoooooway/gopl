package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// 简单的聊天服务。

// 练习 8.12： 使broadcaster能够将arrival事件通知当前所有的客户端。这需要你在clients集合中，以及entering和leaving的channel中记录客户端的名字。

// 练习 8.13： 使聊天服务器能够断开空闲的客户端连接，比如最近五分钟之后没有发送任何消息的那些客户端。
// 提示：可以在其它goroutine中调用conn.Close()来解除Read调用，就像input.Scanner()所做的那样。

// 练习 8.14： 修改聊天服务器的网络协议，这样每一个客户端就可以在entering时提供他们的名字。将消息前缀由之前的网络地址改为这个名字。

// 练习 8.15： 如果一个客户端没有及时地读取数据可能会导致所有的客户端被阻塞。
// 修改broadcaster来跳过一条消息，而不是等待这个客户端一直到其准备好读写。
// 或者为每一个客户端的消息发送channel建立缓冲区，这样大部分的消息便不会被丢掉；broadcaster应该用一个非阻塞的send向这个channel中发消息。

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages

	connTime = make(map[net.Conn]time.Time)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	go broadcaster()

	go Idlechecker()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error while accept: %s\n", err.Error())
			continue
		}

		go handleConn(conn)
	}
}

// 管理所有连接
func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case cli := <-entering:
			// 用户进入
			clients[cli] = true

		case cli := <-leaving:
			// 用户退出
			delete(clients, cli)
			close(cli)

		case msg := <-messages:
			// 处理消息
			for k, _ := range clients {
				select {
				case k <- msg:
					// do nothing
				default:
					// do nothing
				}
			}
		}
	}

}

// 处理连接
func handleConn(conn net.Conn) {
	// 创建一个channel, 代表此连接的客户端, 并通过此channel进行消息的传输
	ch := make(chan string, 50)

	// 启动一个goroutine来后台处理发送到客户端的消息
	go clientWrite(conn, ch)
	input := bufio.NewScanner(conn)

	ch <- "Please enter your username first: "
	input.Scan()
	who := input.Text()
	ch <- "You are: " + who

	entering <- ch

	msg := fmt.Sprintf("%s> %s: %s", time.Now().Format("2006-01-02T15:04:05"), who, "has arrived")
	messages <- msg

	connTime[conn] = time.Now()

	for input.Scan() {
		now := time.Now()
		messages <- fmt.Sprintf("%s> %s: %s", now.Format("2006-01-02T15:04:05"), who, input.Text())
		connTime[conn] = now
	}

	// 处理断开连接
	leaving <- ch
	conn.Close()
	// 通知其他人
	messages <- fmt.Sprintf("%s> %s: %s", time.Now().Format("2006-01-02T15:04:05"), who, "has left")
}

func clientWrite(conn net.Conn, ch <-chan string) {
	for e := range ch {
		// 设置写超时
		conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
		fmt.Fprintln(conn, e)
	}
}

func Idlechecker() {
	tk := time.NewTicker(5 * time.Second)
	for {
		<-tk.C
		for k, v := range connTime {
			if v.Add(5 * time.Minute).Before(time.Now()) {
				k.Close()
				delete(connTime, k)
			}
		}
	}

}
