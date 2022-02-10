package main

import "fmt"

type footballCoach struct{}
type tennisCoach struct{}

func (f footballCoach) giveInstructions() string {
	return "Run 10KM daily"
}
func (t tennisCoach) giveInstructions() string {
	return "Goto gym daily"
}

type coach interface {
	giveInstructions() string
}

func printInstructions(c coach) {
	fmt.Println(c.giveInstructions())
}
func main() {
	f := footballCoach{}
	t := tennisCoach{}
	printInstructions(f)
	printInstructions(t)

}
