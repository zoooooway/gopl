package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("total time is : %d", time.Since(start))

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(io.Discard, res.Body)
	if err != nil {
		ch <- fmt.Sprintf("%s, %s", url, err)
		return
	}
	diff := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%f  %d  %s", diff, nbytes, url)

}
