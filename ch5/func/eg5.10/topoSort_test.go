package function

import (
	"fmt"
	"testing"
)

func BenchmarkTopoSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := topoSort2(prereqs)
		fmt.Printf("%q --> %t\n", r, isVaildTopo(r, prereqs))

	}
	fmt.Println(isVaildTopo([]string{"operating systems", "networks"}, prereqs))
	fmt.Println(isVaildTopo([]string{"computer organization", "data structures", "programming languages"}, prereqs))

}
