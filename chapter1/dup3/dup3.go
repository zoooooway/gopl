package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	paths := os.Args[1:]
	counts := make(map[string]int)

	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			continue
		}

		input := bufio.NewScanner(file)
		for input.Scan() {
			line := input.Text()
			counts[line]++
		}

	}

	for line, count := range counts {
		if count > 1 {
			fmt.Println(count, line)
		}
	}
}
