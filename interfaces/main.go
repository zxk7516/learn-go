package main

import (
	"fmt"
)

type englishBot struct{}
type chineseBot struct{}

type bot interface {
	getGreeting() (int, string)
}

func main() {
	eb := englishBot{}
	cb := chineseBot{}
	printGreeting(eb)
	printGreeting(cb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())

}

func (eb englishBot) getGreeting() (int, string) {
	return 1, "Hi there!"
}

func (cb chineseBot) getGreeting() (int, string) {
	return 2, "你好"
}
