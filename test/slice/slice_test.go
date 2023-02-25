package slice

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestSlice(t *testing.T) {
	// words := []string{"ss", "dd"}
	// var nw []string
	// nw = append(nw, words...)
	// fmt.Println(nw)
	// fmt.Println(words)

	// words[1] = "halo"
	// fmt.Println(nw)
	// fmt.Println(words)

	// ss := make([]string, 10)
	// ns := make([]string, 10)
	// ss[0] = "en"
	// ss[1] = "zh"
	// ns = append(ns, ss...)
	// fmt.Println(ss)
	// fmt.Println(ns)

	// ss[1] = "halo"
	// fmt.Println(ss)
	// fmt.Println(ns)
	// Ensome{}.write()
	var a A
	var w io.Writer = &bytes.Buffer{}
	fmt.Printf("a=(%T, %v)\n", a, a)
	fmt.Printf("w=(%T, %v)\n", w, w)
}

type A struct {
	int
	io.Writer
}

// func (a A) write() {
// 	fmt.Println("test write")
// }
