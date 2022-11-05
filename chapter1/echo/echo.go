package main

import (
	"fmt"
	"os"
)

func main() {
	// s, sep := "", " "

	// for i := 1; i < len(os.Args); i++ {
	// 	s += os.Args[i] + sep
	// }
	// fmt.Print(os.Args[0:])

	// s = "";
	for index, v := range os.Args {
		fmt.Println(index, v)
	}
}
