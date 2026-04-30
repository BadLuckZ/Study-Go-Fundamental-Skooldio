package main

import "fmt"

// Closure function
func oddEven(n int, fn func(n int) bool) {
	fmt.Println(fn(n))
}

func main() {
	a1, b1 := 1, 4

	// Anonymous function
	r1 := func(a int, b int) int {
		return a + b
	}(a1, b1)

	fmt.Printf("%d + %d = %d\n", a1, b1, r1)

	// ===============================

	isOdd := func(a int) bool {
		return a&1 == 1
	}

	isEven := func(a int) bool {
		return a&1 == 0
	}

	oddEven(4, isEven)
	oddEven(4, isOdd)

}
