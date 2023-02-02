package html

import (
	"fmt"
	visit "gopl/ch5/func/eg5.4"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestPrinter(t *testing.T) {
	r := visit.Fetch("http://127.0.0.1:5500/ch5/func/eg5.7/test.html")
	doc, err := html.Parse(r)
	r.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}
