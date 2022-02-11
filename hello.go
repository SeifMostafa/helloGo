package main

import (
	"fmt"
)

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
func main() {
	t := triangle{3.0, 4.0}
	s := square{5}

	printArea(t)
	printArea(s)

}
