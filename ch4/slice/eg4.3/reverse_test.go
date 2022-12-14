package main

import "testing"

func BenchmarkReverseTest(b *testing.B) {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := arr[:]

	for i := 0; i < b.N; i++ {
		reverseNew(&s)
	}
}
