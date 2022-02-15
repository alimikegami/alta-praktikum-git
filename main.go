package main

import "fmt"

// function to greet the user
func greet() {
	fmt.Println("Hi, my name is Alim Ikegami! Nice to meet you!")
	fmt.Println("I Hope you are having a wonderful day!")
	fmt.Println("Have a nice day!")
}

// function to ask question to the user
func askQuestion() {
	fmt.Println("How are you today?")
	fmt.Println("Are you having a good day so far?")
	fmt.Println("Do you want to eat dinner together tonight?")
}

func main() {
	fmt.Println("Hello world")
	greet()
	askQuestion()
}
