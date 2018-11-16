package main

import "fmt"

func main() {
	ix := []int{6, 4, 5, 3, 2, 1}

	unfurl(ix...)
	unfurl2(ix)
}
// Variadic parameters take in zero or more parameters.
//
func unfurl(x ...int) {
	for i, v := range x {
		fmt.Println("Index:", i, "Value:", v)
	}
}

func unfurl2(xi []int) {
	fmt.Println("Inside the Second unfurl func, taking an slice of ints as an argument.")
	for i, v := range xi {
		fmt.Println("Index:", i, "Value:", v)
	}
}
