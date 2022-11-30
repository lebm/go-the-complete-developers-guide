package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://github.com",
		"https://go.dev",
	}

	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	//http.Get() is a blocking call.
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		return
	}
	fmt.Println(link, "is up!")
}
