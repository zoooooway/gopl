package function

import (
	"fmt"
	"testing"
)

func BenchmarkTopoSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("+++++++++++++++++++++++++++++++++++")
		r := topoSort3(prereqs)
		fmt.Printf("%q --> %t\n", r, isVaildTopo(r, prereqs))
		fmt.Println("-----------------------------------")
	}

}
