package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	colors["white"] = "#ffffff"
	colors["black"] = "#000000"

	fmt.Println(colors)

	printMap(colors)

	delete(colors, "white")

	fmt.Println(colors)

	var colors2 map[string]string

	fmt.Println(colors2)

	colors3 := make(map[string]string)

	fmt.Println(colors3)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
