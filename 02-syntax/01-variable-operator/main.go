package main

import "fmt"

func main() {
	// =========== Variables and Operations ============

	var n int = 10 // integer (0 by default)
	fmt.Println(n)

	var f float32 = 0.1 // float (0.0 by default)
	fmt.Println(f)

	var s string = "Hello, World!" // string ("" by default)
	fmt.Println(s)

	n1, n2 := 10, 3
	fmt.Println(n1, n2)

	// Math Operators
	result_sum := n1 + n2
	result_sub := n1 - n2
	result_mul := n1 * n2
	result_div := n1 / n2
	result_mod := n1 % n2
	fmt.Println(result_sum, result_sub, result_mul, result_div, result_mod)

	// Relational Operators (==, !=, >, <, >=, <=)
	fmt.Println(result_sub == result_mul)

	// Logical Operators (&&: and, ||: or, !: not)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Type Conversion
	type char uint8

	var c char = 1 // c is char type (like uint8 but not uint8)
	i := uint8(c)  // i is uint8 type with the same value of c

	fmt.Println("c =", c, " i =", i)

	// =================================
}
