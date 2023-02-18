package reader

import (
	"fmt"
	"io"
	"log"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	r := strings.NewReader("hello world")

	lr := LimitReader(r, 5)
	p := make([]byte, 10)

	n, e := lr.Read(p)
	if e != nil {
		if e != io.EOF {
			log.Fatal(e.Error())
		}
		fmt.Println(e.Error())
	}

	fmt.Println(n)
	fmt.Println(string(p))
}
