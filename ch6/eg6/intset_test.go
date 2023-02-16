package method

import (
	"fmt"
	"testing"
)

func TestIntset(t *testing.T) {
	var it1 IntSet
	var it2 IntSet

	it1.AddAll(3, 5)
	it2.AddAll(5, 1, 2, 9)

	fmt.Println(&it1)
	fmt.Println(&it2)
	re1 := it1.IntersectWith(&it2)
	fmt.Println(re1)
	re2 := it1.DifferenceWith(&it2)
	fmt.Println(re2)
	re3 := it1.SymmetricDifference(&it2)
	fmt.Println(re3)

	fmt.Println(it2.Elems())
	fmt.Println(BIT)
}
