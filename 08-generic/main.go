package main

import "fmt"

type numeric interface {
	int | float64
}

func add[T numeric](a T, b T) T {
	return a + b
}

func main() {
	a1, b1 := 1.4, 4.1
	r1 := add(a1, b1)

	a2, b2 := 1, 4
	r2 := add(a2, b2)

	fmt.Printf("%f + %f = %f\n", a1, b1, r1)
	fmt.Printf("%d + %d = %d\n", a2, b2, r2)

}
