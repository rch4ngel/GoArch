package main

// Defer will defer the function until the last the last statement is ran, panicks or when program calls
// defer.
import "fmt"

func main() {
	// Control flow of this method would be foo and then bar but
	// using defer it will run the deferred function last.
	defer foo()
	bar()
}

func foo() {
	fmt.Println("Foo baby.")
}

func bar() {
	fmt.Println("Bar baby.")
}
