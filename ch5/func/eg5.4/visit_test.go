package visit

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	r := fetch("https://vertx.io/")
	doc, err := html.Parse(r)
	r.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range Visit(nil, "a", doc) {
		fmt.Println(link)
	}
}
