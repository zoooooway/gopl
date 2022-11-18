package main

import "fmt"

func main() {
	var p = new(map[string]int)

	var v = make(map[string]int)

	// 此处会抛出异常 panic: assignment to entry in nil map.
	// 原因：You have to initialize the map using the make function (or a map literal) before you can add any elements.
	// 参见：https://yourbasic.org/golang/gotcha-assignment-entry-nil-map/
	// (*p)["hello"]++

	// *p["hello"]++ // error: invalid operation: cannot index p (variable of type *map[string]int)

	fmt.Println(&p)
	fmt.Println(&(*p))

	v["world"]++
	fmt.Println(v)
	pv := &v
	fmt.Printf("%T\n", pv)
	fmt.Println(&pv)
	fmt.Println(pv)
	fmt.Println(*pv)
}
