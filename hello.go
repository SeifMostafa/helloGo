package main

import (
	"fmt"
	"math"
	"strings"
)

//import "golang.org/x/tour/pic"
///"golang.org/x/tour/wc"

func main() {
	fmt.Println("Hello, World!")
	//pic.Show(Pic)
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))    // use 3 & 4 from compute in hypot function
	fmt.Println(compute(math.Pow)) // 3 to power 4
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func Pic(dx, dy int) [][]uint8 {
	z := make([][]uint8, dy)
	for i := range z {
		z[i] = make([]uint8, dx)
		for j := range z[i] {
			z[i][j] = uint8(i * j)
		}
	}
	return z
}
func WordCount(s string) map[string]int {
	v := strings.Fields(s)
	mapy := make(map[string]int)
	for _, word := range v {
		_, ok := mapy[word]
		if ok {
			mapy[word] = mapy[word] + 1
		} else {
			mapy[word] = 1
		}
	}
	return mapy
}
