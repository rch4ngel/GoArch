package main

import "fmt"

// Closure is the technique of "closing" the scope around a variable.
// Closure is a very zen aspect of developing.
// Always work on limiting scope.

var x int // The scope of x is the whole package.

func main() {
	var y int // The scope of y is in the main function
	y = 1
	fmt.Println(y)
	{
		var z int // The scope of z is limited to this code block...
		z = 420
		fmt.Println(z)
	}

	printX()

	// Calling the incrementer
	i := incrementer()
	fmt.Println("Calling incrementer:", i())
	fmt.Println("Calling incrementer:", i())
	fmt.Println("Calling incrementer:", i())
	fmt.Println("Calling incrementer:", i())
}

func printX() {
	x = 69
	fmt.Println(x)
}
