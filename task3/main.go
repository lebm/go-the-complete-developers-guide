package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		file, err := os.Open(arg)
		if err != nil {
			fmt.Printf("Error opening %v %v", arg, err)
			continue
		}
		io.Copy(os.Stdout, file)
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
