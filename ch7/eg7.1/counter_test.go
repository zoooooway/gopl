package counter

import (
	"fmt"
	"testing"
)

func TestCounter(t *testing.T) {
	var wc WordCount
	var lc LineCounter
	str := `hello i am lily
	hi lily good to see you
	yeah nice to see you too`
	count, e := wc.write([]byte(str))
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println(count)
	count, e = wc.write([]byte(str))
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println(count)
	fmt.Println(wc)

	fmt.Println("--------------------------")

	count, e = lc.write([]byte(str))
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println(count)
	count, e = lc.write([]byte(str))
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println(count)
	fmt.Println(lc)
}
