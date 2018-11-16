package main

import "fmt"

// Package Level Variables
var xx int
var yy string
var zz bool

type age int

var a age = 33

func main() {
	// Identifiers below:
	x := 42
	y := "Batman"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(xx, yy, zz)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

	s := fmt.Sprintf("%v\t%b\t%#x\t", x, x, x)
	fmt.Println(s)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", int(a))
}
