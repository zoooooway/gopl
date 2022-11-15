package main

import (
	"fmt"
	"os"
	"strconv"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	x, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Printf("%d has %d bits that are 1", x, PopCount(x))
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	// byte是uint8的别名，因此byte(x)将舍弃高位
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
