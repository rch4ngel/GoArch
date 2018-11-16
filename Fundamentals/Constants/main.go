package main

import "fmt"

// Dealing with constants. Constants can be declared with types
// or without a type. When it is not typed it is an untyped constant,
// known as a constant of a kind as well.

const (
	s int = 9
	b     = 9.9
	t     = "Look at me"
)

func main() {
	fmt.Println("Constant values are ", s, b, t)
	fmt.Printf("The types are %T, %T, %T\n", s, b, t)
}
