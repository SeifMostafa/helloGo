package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

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

	response, err := http.Get("http://www.google.com")
	fmt.Printf("%+v , %+v", response, err)

	// bs := make([]byte, 99999)
	// response.Body.Read(bs)
	// fmt.Println(string(bs))

	io.Copy(os.Stdout, response.Body)
}
