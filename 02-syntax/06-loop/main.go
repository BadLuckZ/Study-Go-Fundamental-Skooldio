package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

	var j int = 2
	for ; j < 20; j += 2 {
		if j < 14 {
			fmt.Printf("j = %d\n", j)
		} else if j == 14 {
			fmt.Println("j = 14 so skip it...")
			continue
		} else {
			fmt.Println("j is more than 14 so break the loop...")
			break
		}
	}
}
