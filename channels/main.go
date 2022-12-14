package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://github.com",
		"https://go.dev",
		"https://httpbin.org/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	//http.Get() is a blocking call.
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s Might be down I think\n", link)
		c <- link
		return
	}
	fmt.Printf("Yep, %s its up.\n", link)
	c <- link
}
