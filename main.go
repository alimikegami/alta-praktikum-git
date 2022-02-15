package main

import "fmt"

// function to greet the user
func greet() {
	fmt.Println("Hi, my name is Alim Ikegami! Nice to meet you!")
	fmt.Println("I Hope you are having a wonderful day!")
	fmt.Println("Have a nice day!")
}

func askQuestion() {
	fmt.Println("How are you today?")
}

func main() {
	fmt.Println("Hello world")
	greet()
	askQuestion()
}
