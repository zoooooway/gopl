package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	paths := os.Args[1:]
	for _, path := range paths {
		counts := make(map[string]int)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			continue
		}
		lines := strings.Split(string(data), "\n")

		for _, line := range lines {
			counts[line]++
		}
		for line, count := range counts {
			if count > 1 {
				fmt.Printf("%s:\t", path)
				fmt.Println(count, line)
			}
		}
	}

}
