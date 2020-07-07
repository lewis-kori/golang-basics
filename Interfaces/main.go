package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// VERY custom logic for an english greeting
	return "Hello There mate!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for an spanish greeting
	return "Aloha pendejo!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
