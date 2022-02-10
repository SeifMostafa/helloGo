package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//import "golang.org/x/tour/pic"
///"golang.org/x/tour/wc"
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
} // function closure

func main() {
	// fmt.Println("Hello, World!")
	// //pic.Show(Pic)
	// hypot := func(x, y float64) float64 {
	// 	return math.Sqrt(x*x + y*y)
	// }
	// fmt.Println(hypot(5, 12))

	// fmt.Println(compute(hypot))    // use 3 & 4 from compute in hypot function
	// fmt.Println(compute(math.Pow)) // 3 to power 4

	// cards := deck{"Ok", "No", "Yes"}

	// fmt.Println(cards)

	// for i, card := range cards {
	// 	println(i, card)
	// }
	// c := color("Red")

	// fmt.Println(c.describe("is an awesome color"))
	beautifulQuadruples(1150, 1547, 853, 423)
}

type color string

func (c color) describe(description string) string {
	return string(c) + " " + description
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func timeConversion(s string) string {
	timeSection := s[len(s)-2:]
	result := s[:len(s)-2]

	if timeSection == "AM" {
		if s[:2] == "12" {
			result = "00" + result[2:]
		}
	} else {
		if s[:2] != "12" {
			hoursIn24, err := strconv.Atoi(s[:2])
			if err == nil {
				hoursIn24 += 12
				result = fmt.Sprint(hoursIn24) + result[2:]
			}
		}
	}
	return result
}
func beautifulQuadruples(a int32, b int32, c int32, d int32) int32 {

	stored := []string{}
	for w := int32(1); w <= a; w++ {
		for x := int32(1); x <= b; x++ {
			for y := int32(1); y <= c; y++ {
				for z := int32(1); z <= d; z++ {
					if w^x^y^z != 0 {
						s := []int32{w, x, y, z}
						sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
						ss := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(s)), ""), "[]")

						if !contains(stored, ss) {
							stored = append(stored, ss)
						}
					}
				}
			}
		}
	}
	return int32(len(stored))
}
func contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
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
