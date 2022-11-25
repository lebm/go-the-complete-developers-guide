package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	//resp, err := http.Get("https://google.com")
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// contentBody := make([]byte, 99999)

	// n, err := resp.Body.Read(contentBody)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(resp)
	// fmt.Println(n)

	// fmt.Println(string(contentBody))

	io.Copy(os.Stdout, resp.Body)
}
