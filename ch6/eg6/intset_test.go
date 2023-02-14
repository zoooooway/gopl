package method

import (
	"fmt"
	"testing"
)

func TestIntset(t *testing.T) {
	var it1 IntSet
	var it2 IntSet

	it1.AddAll(1, 2, 3)
	it2.AddAll(2, 9, 1)

	fmt.Println(&it1)
	fmt.Println(&it2)
	re := it1.IntersectWith(&it2)
	fmt.Println(re)

}
