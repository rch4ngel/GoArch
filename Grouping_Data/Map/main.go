package main

import "fmt"

// Here is an example of a map. Maps are unordered lists.
// Maps are very fast for lookups.
// When working with maps, if a key has no value it will
// return the the empty type. The way to check for that is explained below.

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	// Check and make sure that the key has a value with the comma ok idiom.
	v, ok := m["Jessie"]
	fmt.Println(v)
	fmt.Println(ok)

	// Lets get idiomatic!
	if v, ok := m["Jessie"]; ok{
		println(v)
	}

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("Inside the bad ass idiomatic statment. Moneypenny is:", v, "years old!")
	}

	fmt.Println(m)
}
