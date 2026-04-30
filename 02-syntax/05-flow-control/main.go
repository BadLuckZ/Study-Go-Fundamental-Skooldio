package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Scanln(&age)

	// If-Else
	if age > 60 {
		fmt.Printf("Your age = %d so you are too old to drink alcohol.\n", age)
	} else if age > 20 {
		fmt.Printf("Your age = %d so you can legally drink alcohol.\n", age)
	} else {
		fmt.Printf("Your age = %d so you can't drink alcohol.\n", age)
	}

	// Switch
	switch {
	case age < 20:
		fmt.Println("You should wait...")
	case age < 60:
		fmt.Println("It's up to you man...")
	default:
		fmt.Println("Please rest...")
	}
}
