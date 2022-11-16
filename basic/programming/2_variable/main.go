package main

import "fmt"

func main() {
	var name string   // Define variable with type string
	fmt.Println(name) // Print variable -> gonna print zero/default value of string which is empty string
	name = "John Doe" // Assign value into variable
	fmt.Println(name) // Print Variable

	greeting := "Hello World" // Define variable and give it a value (:= Short Variable Declaration, sometime people call it shorthand operator)
	fmt.Println(greeting)     // Print Variable
}
