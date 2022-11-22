package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards.txt")
	cards2 := newDeckFromFile("my_cards.txt")
	fmt.Println(cards2)
	cards2.shuffle()
	cards2.print()
}
