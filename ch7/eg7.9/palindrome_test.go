package palindrome

import (
	"fmt"
	"sort"
	"testing"
)

func TestPalindrome(t *testing.T) {
	list := sort.IntSlice{1, 2, 3, 4, 3, 2, 1}
	fmt.Println(IsPalindrome(list))
}
