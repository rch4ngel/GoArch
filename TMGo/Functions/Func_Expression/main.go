package main

import "fmt"

// In go functions are first class citizens. Which
// means it can be used like any other type.

func main() {
	f := func() {
		fmt.Println("Executing my first function expression. ")
	}

	g := func(s string) {
		fmt.Println("The number one samurai ever is:", s)
	}

	f()
	g("Jubei")
}
