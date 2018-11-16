package main

import "fmt"

// Here we will be creating directional channels for our programs.
// Creating directional channels are useful for readability  when
// wanting to say a certain function or method only receives channels,
// or sends channels.

func main() {
	// Created a buffer to hold two channels.
	c := make(chan int, 2)

	// These are bidirectional channels. We send a channel
	// and Println receives the channel.
	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\n", c)

	fmt.Println("-------")
	// Unlike above the code below defines directional channels

	// c2 channel is only a send channel.
	cs := make(chan <- int, 1) // send channel
	fmt.Printf("%T\n",  cs)
	// fmt.Println(<-c2) this code will not work since we are receiving a channel for println.

	// Creating a receive only channel.
	cr := make(<- chan int, 1) // receive channel
	fmt.Printf("%T\n", cr)

}