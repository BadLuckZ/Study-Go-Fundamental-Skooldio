package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func hello(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	var num1, num2 int = 2, 3
	var num_sum = add(num1, num2)

	fmt.Println(num1, "+", num2, "=", num_sum)

	name := "Jack"
	hello(name)
}
