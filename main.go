package main

import "fmt"

func greet() {
	fmt.Println("Hi, my name is Alim Ikegami! Nice to meet you!")
}

// function to ask question to the user
func askQuestion() {
	fmt.Println("How are you today?")
}

func main() {
	fmt.Println("Hello world")
	greet()
	askQuestion()
}
