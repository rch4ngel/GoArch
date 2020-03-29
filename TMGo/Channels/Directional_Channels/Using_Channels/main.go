package main

import "fmt"

// Here we are giving an example on how to use the directional channels when creating functions,
// or methods.

func main() {
	c := make(chan int)
	go sendChannel(c) // This go routine will go off and set the channel to 42.
	receiveChannel(c) // This will block the program until it pulls the value off of the go routine.
}

func sendChannel(c chan<- int) {
	c <- 42
}

func receiveChannel(c <-chan int) {
	fmt.Println(<-c)
}
