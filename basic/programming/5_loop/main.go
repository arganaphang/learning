package main

import "fmt"

func main() {
	// ---- FOR RANGE
	fruits := []string{"Mango", "Banana", "Pineapple", "Watermelon"} // Define variable with type array of string and give it a bunch of fruit names
	for _, fruit := range fruits {
		fmt.Println(fruit) // Print each item of array fruits
	}

	// ---- FOR
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // Define variable with type array of int and give it a bunch of numbers
	for i := 0; i < len(numbers)-1; i++ {
		fmt.Println(numbers[i]) // Print each item of array numbers
	}

	// ---- FOR CONDITIONAL / WHILE LOOP
	i := 0      // Define variable i = 0
	for i < 5 { // if i less than 5, this loop execute job inside it
		fmt.Println(i) // Print i
		i++            // Increment i each loop
	}
}
