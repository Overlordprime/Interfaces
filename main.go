package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	sb := spanishBot{}
	eb := englishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//Logic for English bot
	return "Hello"
}

func (sb spanishBot) getGreeting() string {
	//Logic for Spanish bot
	return "Hola"
}
