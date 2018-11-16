package main

import "fmt"

// Here we will append to a slice
func main() {
	s := []int{5 ,2, 5, 3, 3, 2}
	for i, v := range s {
		fmt.Println("Index: ", i, "\tValue: ", v)
	}

	s = append(s, 44, 22, 44, 11)

	for i, v := range s {
		fmt.Println("Index: ", i, "\tValue: ", v)
	}

	// How do we append on slice to another slice?
	s2 := []int{ 33, 44, 22, 11}

	// The syntax below with the ... is a variadic parameter.
	// Variadic parameters allow for n amount of variables.
	// another term used along side this is unfurling.
	s = append(s, s2...)
	for i, v := range s {
		fmt.Println("Index: ", i, "\tValue: ", v)
	}
}
