package main

import (
	"fmt"
	"net/http"
	"sync"
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

	var wg sync.WaitGroup

	for _, link := range links {
		wg.Add(1)
		l := link
		go func() {
			defer wg.Done()
			checkLink(l)
		}()
	}

	wg.Wait()
}

func checkLink(link string) {
	//http.Get() is a blocking call.
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		return
	}
	fmt.Println(link, "is up!")
	//	fmt.Println(resp)
}
