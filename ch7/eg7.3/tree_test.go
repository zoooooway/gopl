package tree

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var tr tree
	tr.value = 1
	add(&tr, 2)
	add(&tr, 3)
	add(&tr, 4)

	fmt.Println(tr.String())
}
