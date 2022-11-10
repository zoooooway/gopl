package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int

var mu sync.Mutex

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprint(w, r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	// 如果在此处使当前例程休眠，同时访问除/count外的任意路径来触发handler处理，则handler会阻塞等待此处休眠结束后执行mu.Unlock()结束
	// time.Sleep(6 * 1000 * 1000 * 1000)
	fmt.Fprint(w, count)
	mu.Unlock()
}
