package main

import "fmt"

// To declare a variable outide of a function use var
var o int = 42

// The func main is the entry point of the program
// when the last statement is finished in the main
// function the program will end.
// This process is called control flow.
func main() {
	// To declare and assign a value within a function
	// use the short declaration operator.
	v := 2

	fmt.Println(o, v)
}
