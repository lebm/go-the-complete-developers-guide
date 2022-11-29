package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

// As logWriter implements Write, it implements the Writer interface.
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Escrevi %v bytes.\n", len(bs))
	return len(bs), nil
}

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

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}
