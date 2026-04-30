package main

import "fmt"

func main() {
	m := map[string]string{
		"b": "banana",
		"c": "coconut",
	}
	// {"b": "banana", "c": "coconut"}

	m["a"] = "apple"
	// {"b": "banana", "c": "coconut", "a": "apple"}

	for k, v := range m {
		fmt.Printf("Key = %s -> Value = %s\n", k, v)
	}

}
