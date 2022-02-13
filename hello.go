package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://www.google.com",
	}
	c := make(chan string)
	for _, url := range urls {
		go checkURL(url, c)
	}
	fmt.Println(<-c)
}
func checkURL(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR!", err)
		c <- "error"
		return
	}
	//io.Copy(os.Stdout, response.Body)
	fmt.Println("OK.")
	c <- "ok"

}
