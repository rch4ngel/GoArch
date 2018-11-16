package main

import "fmt"

// A send channel
func main() {
	cs := make(chan<- int, 1)
	fmt.Printf("%T\n", cs)
}
