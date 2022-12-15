package main

import (
	"fmt"
	"os"
)

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	empty := map[string]int(nil)
	zero := map[string]bool{}
	zero["hash"] = true
	fmt.Println(zero["hash"])
	fmt.Println(zero["dosh"])

	fmt.Println(ages == nil)
	fmt.Println(empty == nil)
	fmt.Println(zero == nil)

	fmt.Println(fmt.Sprintf("%q", os.Args))

	// fmt.Println(ages == empty) // invalid operation: cannot compare ages == empty (map can only be compared to nil)
	empty = nil
	// fmt.Println(ages == empty) // invalid operation: cannot compare ages == empty (map can only be compared to nil)

}
