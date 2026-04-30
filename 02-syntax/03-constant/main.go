package main

import "fmt"

func main() {
	const DEFAULT_VALUE_NUMBER = 1 // can't change DEFAULT_VALUE after this
	// const type is "untype"

	const DEFAULT_VALUE_STRING = "Go"
	// can't change DEFAULT_VALUE_STRING after this

	fmt.Println(DEFAULT_VALUE_NUMBER, DEFAULT_VALUE_STRING)

	// iota
	const (
		SUNDAY    = iota // 0
		MONDAY           // 1
		TUESDAY          // 2
		WEDNESDAY        // 3
		THURSDAY         // 4
		FRIDAY           // 5
		SATURDAY         // 6
	)

	fmt.Println(SUNDAY, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY)

	const (
		BYTE = 1 << (iota * 10) // 1
		KB                      // 1024
		MB                      // 1048576
	)

	fmt.Println(BYTE, KB, MB)
}
