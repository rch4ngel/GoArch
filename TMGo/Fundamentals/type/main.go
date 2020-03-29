package main

import "fmt"

type age int

var a age

func main() {
	i := 55
	a = 29
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", a)
}
