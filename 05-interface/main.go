package main

import "fmt"

type dog string

func (d dog) Sound() string {
	return "Woof!"
}

type cat string

func (c cat) Sound() string {
	return "Meow!"
}

type animal interface {
	Sound() string
}

func playSound(a animal) string {
	return a.Sound()
}

func main() {
	var a any

	a = 10
	fmt.Printf("%T %v\n", a, a) // int 10

	a = "ten"
	fmt.Printf("%T %v\n", a, a) // string ten

	// ======================================

	var d dog
	var c cat

	fmt.Printf("%T's sound: %s\n", d, playSound(d))
	fmt.Printf("%T's sound: %s\n", c, playSound(c))
}
