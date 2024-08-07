package main

import "fmt"

const (
	genGrievous = "Grevious"
	kenobi      = "Obi Wan"

	firstGreeting  = "General Kenobi"
	secondGreeting = "Hello There"
)

func main() {
	fmt.Println(HelloThere())
}

func HelloThere() string {
	return fmt.Sprintf("%s: %s, %s: %s", genGrievous, firstGreeting, kenobi, secondGreeting)
}
