package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]

	switch args {
	case "-384":
		fmt.Println(sha512.Sum384([]byte(os.Args[len(os.Args)-1])))
	case "-512":
		fmt.Println(sha512.Sum512([]byte(os.Args[len(os.Args)-1])))
	default:
		fmt.Println(sha256.Sum256([]byte(os.Args[len(os.Args)-1])))
	}

}
