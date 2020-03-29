package main

import "fmt"

// Conversion is a way to convert a type to another type.

type original int

func main() {
	// Here o is declared and assigned of type original with the underlying type
	// being int.
	var o original
	o = 42
	fmt.Printf("o is of type %T\n", o)

	// Since o has an underlying type of int, it can be converted back to int.
	// In this case o is converted to int and assigned to i which is of type int.
	var i int
	i = int(o)
	fmt.Printf("o was converted to int and assigned to i. i is of type %T\n", i)

}
