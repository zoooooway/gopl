package main

import "fmt"

func counter(out chan<- int) {
	fmt.Printf("c out: %T, %v\n", out, out)
	fmt.Printf("c equal naturals?: %t\n", out == naturals)
	for x := 0; x < 1; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	fmt.Printf("s out: %T, %v\n", out, out)
	fmt.Printf("s out equal squares?: %t\n", out == squares)
	fmt.Printf("s in: %T, %v\n", in, in)
	fmt.Printf("s in equal naturals?: %t\n", in == naturals)
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	fmt.Printf("p in: %T, %v\n", in, in)
	fmt.Printf("p in equal squares?: %t\n", in == squares)
	for v := range in {
		fmt.Println(v)
	}
}

var naturals = make(chan int)
var squares = make(chan int)

func main() {
	fmt.Printf("naturals: %T, %v\n", naturals, naturals)
	fmt.Printf("squares: %T, %v\n", squares, squares)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
