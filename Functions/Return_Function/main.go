package main

import "fmt"

// Since functions are first class citizens...
// They can also be returned.
// Why would you ever do this??
// well sometimes you return functions. Good example
// for this is the HandleFunc function. 

func main() {
	s := foo()
	fmt.Println(s)

	b := bar()
	i := b()

	fmt.Println(i)

	// A cleaner way of calling the return function
	w := wine()
	fmt.Println(w())

	// Or even a cleaner way...
	fmt.Println(wine()())
}

func foo() string {
	s := "Ryu"
	return s
}

// Here is a func that returns a func that returns an int.
func bar() func() int {
	return func() int {
		return 420
	}
}

func wine() func() string {
	return func() string {
		return "Clark Kent"
	}
}
