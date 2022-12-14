package main

import "testing"

func BenchmarkReverseTest(b *testing.B) {

	x := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		rotate(x[:], i%10, true)
	}
}
