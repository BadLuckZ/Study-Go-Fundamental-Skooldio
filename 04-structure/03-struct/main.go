package main

import "fmt"

type rect struct {
	width  float64
	length float64
}

func (r *rect) setWidth(width float64) {
	r.width = width
}

func (r *rect) setLength(length float64) {
	r.length = length
}

func (r rect) findArea() float64 {
	return r.width * r.length
}

func main() {
	r := rect{
		width:  3,
		length: 4,
	}

	area := r.width * r.length
	fmt.Printf("Rectangle area with width = %f and length = %f is %f\n", r.width, r.length, area)

	r.setLength(6)
	r.setWidth(8)

	area = r.findArea()

	fmt.Printf("Rectangle area with width = %f and length = %f is %f\n", r.width, r.length, area)
}
