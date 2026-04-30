package main

import "fmt"

func double(n *int) {
	*n *= 2
}

func main() {
	n := 14
	// [14]n

	var p *int = &n
	// p is the pointer pointing to n's address
	// p --> [14]n

	*p += 1
	// p --> [15]n

	fmt.Println(n)
	// n = 15

	double(p)
	// p --> [30]n

	fmt.Println(n)
	// n = 30
}
