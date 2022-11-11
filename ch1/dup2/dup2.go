package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	files := os.Args[1:]
	counts := make(map[string]int)
	if len(files) == 0 {

		std := os.Stdin
		countLines(std, counts)

	} else {
		for _, path := range files {
			file, err := os.Open(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error opening")
				continue
			}
			countLines(file, counts)
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Println(line, count)
		}
	}

}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		counts[line]++
	}
}
