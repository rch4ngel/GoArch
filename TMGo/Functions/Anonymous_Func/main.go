package main

import "fmt"

// Anonymous functions also know as self calling functions are
// written in the form func(parameters) {body} (arguments)
func main() {
	func(x int) {
		fmt.Println("Hello I am an anonymous function")
		fmt.Println(x, "was passed into this anonymous function")
	}(42)
}
