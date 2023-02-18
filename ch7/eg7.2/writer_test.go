package write

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {

	w, count := CountingWriter(os.Stdout)
	n, e := w.Write([]byte("hello word!"))
	if e != nil {
		log.Fatal(e.Error())
	}
	fmt.Println(n)
	fmt.Println(*count)

	n, e = w.Write([]byte("hello word!"))
	if e != nil {
		log.Fatal(e.Error())
	}
	fmt.Println(n)
	fmt.Println(*count)

}
