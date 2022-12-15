package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	var words []string
	var freq map[string]int
	for input.Scan() {
		words = append(words, input.Text())
	}
	freq = wordfreq(words)
	for k, v := range freq {
		fmt.Printf("%s  %d\n", k, v)
	}
}

// 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。
// 在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。
func wordfreq(s []string) map[string]int {
	freq := make(map[string]int)
	for _, v := range s {
		freq[v]++
	}
	return freq
}
