package main

import "fmt"

func main() {
	// -- Arithmetic Operators
	a := 2 + 4     // Addition(+) Arithmetic Operator
	fmt.Println(a) // Print Variable

	// -- Relational Operators
	age := 16
	canDrink := age > 21 // Compare age if age GREATER THAN -> value must be true
	if canDrink {
		fmt.Println("Yep you can drink")
	} else {
		fmt.Println("No, not yet")
	}
}
