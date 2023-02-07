package html

import (
	"fmt"
	"os"
	"testing"

	visit "gopl/ch5/func/eg5.4"

	"golang.org/x/net/html"
)

func TestHTML(t *testing.T) {
	var depth int

	r := visit.Fetch("http://127.0.0.1:5500/ch5/func/eg5.7/test.html")
	doc, err := html.Parse(r)
	r.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	ns := ElementsByIds(doc, "haha1", "haha2")
	for _, v := range ns {
		fmt.Printf("%*s<%s %s>\n", depth*2, "", v.Data, toString(v.Attr))
	}

}
