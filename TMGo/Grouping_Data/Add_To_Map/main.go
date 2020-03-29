package main

import "fmt"

func main() {
	m := map[string]int{
		"Jason": 44,
		"Mike":  25,
	}

	m["Lilly"] = 44

	for k, v := range m {
		fmt.Println("Key:", k, "\tValue:", v)
	}
}
