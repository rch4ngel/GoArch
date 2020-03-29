package main

import "fmt"

// This file is Important in the fact that it will delve a little deeper
// into slices.
// By in large it is good practice to initialize slices using a composite
// literal.

// Slices run on top of an array. When you know the size of the array, it is good to
// use make. Make will create the underlying array with that size. This will cut down
// on the processing power of resizing the underlying array. BOOOM.

func main() {
	// So this will create the underlying array with the size of 100.
	// It will then create the overlying map with 10 elements. :)
	s := make([]int, 10, 100)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	// Here is an example of the underlying functionality.
	// Created below is a slice made with the underlying array
	// size being 10. We will fill up the ten values. We will then
	// go one more size above the underlying capacity. This will
	// require the array to copy its contents, destroy original array, create a new array
	// of double the size, copy old data into new array. Note append is the function that
	// has to be used to do this.

	s2 := make([]int, 9, 10)
	fmt.Println(s2)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s2[0] = 1
	s2[8] = 221
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	s2 = append(s2, 420, 69)
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
}
