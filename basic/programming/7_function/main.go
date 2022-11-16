package main

import "fmt"

func main() {
	greeting()                // Just Call function
	add := calculateAdd(2, 4) // Call function and store return value into variable
	fmt.Println(add)          // Print variable
}

// Void Function (function with no return value)
func greeting() {
	fmt.Println("Hello World")
}

// Non-Void Function (function with return value)
func calculateAdd(a, b int) int {
	return a + b
}
