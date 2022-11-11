package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("%d\n", res.StatusCode)
		io.Copy(os.Stdout, res.Body)
		res.Body.Close()
	}
}
