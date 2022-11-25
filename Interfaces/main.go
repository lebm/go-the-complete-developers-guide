package main

import "fmt"

// The "bot" type is an interface.
// It defines the "getGreeting()" function.
// It means that, ever type that defines its own getGreeting() function will also be a "bot".
// These type implements implicitly this interface. They can be used everywhere a bot would be used.
type bot interface {
	getGreeting() string
}

// printGreeting has a bot as a parameter, and a bot implements "getGreeting()",
// so printGreeting can call getGreeting on its argument at runtime.
// Every type that implements getGreeting is also a bot, so values of the type can be passed to printGreeting.
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Both englishBot and spanishBot implements getGreeting, so they are also "bot"s.
// Values of thess types can be passed to printGreeting()
type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

// func printGreeting(en englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
