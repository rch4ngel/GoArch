package main

import "fmt"

func main() {
	x := 69
	y := 420
	z := float32(y) / float32(x)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
