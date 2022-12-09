package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	for _, v := range args {
		if v == "-256" {
			fmt.Println(sha256.Sum256([]byte(args[0])))
			break
		}

		if v == "-384" {
			fmt.Println(sha512.Sum384([]byte(args[0])))
			break
		}

		if v == "-512" {
			fmt.Println(sha512.Sum512([]byte(args[0])))
			break
		}
	}

}
