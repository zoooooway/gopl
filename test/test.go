package main

import (
	"fmt"
	"time"
)

func main() {
	// go spinner(100 * time.Millisecond)
	// const n = 45
	// fibN := fib(n) // slow
	// fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

	ch := make(chan int, 2)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("--")
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
			fmt.Println("++")
		}
	}
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `_\|/` {
			fmt.Printf("%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
