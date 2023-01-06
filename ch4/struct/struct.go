package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

type Point struct {
	X, Y int
}

type god struct {
	X, Y int
}

type Circle struct {
	Point
	god
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	p, err := 1, 2

	if err := 4 * 3; err > 3 {
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Println(err)

	p1 := Point{1, 2}

	p2 := Point{X: 1}
	fmt.Println(p1) // {1 2}
	fmt.Println(p2) // {1 0}

	w := Wheel{
		Circle: Circle{
			Point:  Point{X: 1, Y: 2},
			god:    god{X: 3, Y: 4},
			Radius: 6,
		},
		Spokes: 2,
	}

	fmt.Println(w.Point.X)
	fmt.Println(w.god.X)
	// fmt.Println(w.X) // error: ambiguous selector w.X
}
