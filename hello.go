package main

import "fmt"

// func main() {
// 	sliceOfInts := makeRange(0, 10)
// 	for i := 0; i < len(sliceOfInts); i++ {
// 		if sliceOfInts[i]%2 == 0 {
// 			fmt.Println("even")
// 		} else {
// 			fmt.Println("odd")
// 		}
// 	}
// }

func main() {
	Seif := person{"Seif", "Mostafa"}
	//println(Seif.firstName)
	fmt.Println(Seif)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
