package main

import (
	"fmt"
	"net/http"
	"time"
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
		time.Sleep(time.Second * 5) // to wait 5 seconds before checking website again
		go checkURL(<-c, c)
	}
}
func checkURL(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, " is down!")
		c <- url
		return
	}
	//io.Copy(os.Stdout, response.Body)
	fmt.Println(url, "is up & running!")
	c <- url

}
