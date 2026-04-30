package main

import "fmt"

func decision(a any) {
	v, ok := a.(int)
	if ok {
		fmt.Println("a is int and a =", v)
		var n int = v
		fmt.Println(n)
		return
	} else {
		fmt.Println("a is not int")
	}
}

func main() {
	var a any
	a = 10
	decision(a)

	a = "10"
	decision(a)
}
