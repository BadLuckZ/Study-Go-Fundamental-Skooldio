package greet

import "testing"

func TestGreetName(t *testing.T) {
	given := "Jack"
	expect := "Hi, Jack"

	result := greet(given)

	if expect != result {
		t.Errorf("given %q, expect %q, actual %q", given, expect, result)
	}
}
