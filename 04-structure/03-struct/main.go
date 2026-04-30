package main

import "fmt"

type rect struct {
	width  float64
	length float64
}

func main() {
	r := rect{
		width:  3,
		length: 4,
	}

	area := r.width * r.length
	fmt.Printf("Rectangle area with width = %f and length = %f is %f\n", r.width, r.length, area)
}
