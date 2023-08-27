package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	var links []string

	if len(os.Args) > 1 {
		links = os.Args[1:]
	} else {
		links = []string{
			"http://google.com",
			"http://facebook.com",
			"http://stackoverflow.com",
			"http://golang.org",
			"http://amazon.com",
		}
	}

	c := make(chan string)

	for _, link := range links {
		go checkStatus(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkStatus(link, c)
		}(l)
	}
}

func checkStatus(link string, c chan string) {
	status, err := checkLink(link, c)
	if status == "down" {
		fmt.Println(link, "Might be down!", err)
	} else {
		fmt.Println(link, "Is up!")
	}
	c <- link
}

func checkLink(link string, c chan string) (string, error) {
	_, err := http.Get(link)
	if err != nil {
		return "down", err
	}
	return "up", nil
}