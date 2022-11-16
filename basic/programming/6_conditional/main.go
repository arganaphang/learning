package main

import "fmt"

func main() {
	isRegister := false // Define variable and give it boolean value
	if isRegister {     // Conditional check must be boolean value
		fmt.Println("Not register yet") // Print if conditional passed
	} else {
		fmt.Println("Already registered") // Print if consitional didn't pass
	}
}
