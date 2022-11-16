package main

import "fmt"

func main() {
	fruits := []string{"Mango", "Banana", "Pineapple", "Watermelon"} // Define variable with type array of string and initiate value into variable
	fmt.Println(fruits)                                              // Print all item of array
	fmt.Println(fruits[3])                                           // Print Spesific Item Array
	fruits[2] = "Banana Update"                                      // Update Item Array
	fmt.Println(fruits)                                              // Print all item of array
	fruits = append(fruits, "Gomu Gomu no Mi")                       // (😉) Add item into array
	fmt.Println(fruits)                                              // Print all item of array
	fruits = append(fruits[:0], fruits[0+1:]...)                     // Delete spesific item of array (Mango)
	fmt.Println(fruits)                                              // Print all item of array
}
