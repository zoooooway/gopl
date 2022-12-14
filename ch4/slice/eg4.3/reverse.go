package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	// reverseNew(&arr[:]) // invalid operation: cannot take address of arr[:] (value of type []int)
	s := arr[:]
	// reverseNew(&s)
	rotateRight(s)
}

// 重写reverse函数，使用数组指针代替slice
func reverseNew(ptr *[]int) {

	for i, j := 0, len(*ptr)-1; i < j; i, j = i+1, j-1 {
		(*ptr)[i], (*ptr)[j] = (*ptr)[j], (*ptr)[i]
	}
}

func rotateLeft(s []int) {
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[3 4 5 1 2]"
}

func rotateRight(s []int) {
	// Rotate s left by two positions.
	reverse(s)
	reverse(s[:2])
	reverse(s[2:])
	fmt.Println(s) // "[2 3 4 5 0 1]"
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
