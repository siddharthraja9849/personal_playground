package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type engBot struct {
}

type spBot struct {
}

func (engBot) getGreeting() string {
	// logic 1
	return "HELLO WORLD!!"
}

func (spBot) getGreeting() string {
	// logic 2
	return "HOLA WORLD!!"
}

func printGreeting(bot Bot) {
	fmt.Println(bot.getGreeting())
}

func main() {
	eBot := engBot{}
	sBot := spBot{}
	printGreeting(eBot)
	printGreeting(sBot)
}
