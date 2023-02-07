package math

import (
	"fmt"
	"testing"
)

func TestMath(t *testing.T) {
	fmt.Println(quicksort([]int{0, 13, 5, 6, 9, 33, 23}))
}
