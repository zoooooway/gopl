package main

import "fmt"

func main() {
	var s int
	x := -4
	switch {
	case x > 0:
		s = 1
		// default case should be first or last in switch statement.
		// but it still works correctly
	default:
		s = 0
	case x < 0:
		s = -1
	}
	fmt.Println(s)
}
