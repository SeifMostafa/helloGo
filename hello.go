package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.musala.com",
	}
	c := make(chan string)
	for _, url := range urls {
		go checkURL(url, c)
	}
	for {
		go checkURL(<-c, c)
	}
}
func checkURL(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR")
		c <- url
		return
	}
	//io.Copy(os.Stdout, response.Body)
	fmt.Println("OK")
	c <- url

}
