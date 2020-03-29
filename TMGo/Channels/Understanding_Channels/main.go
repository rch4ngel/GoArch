package main

import "fmt"

// Understanding channels. Here we will talk about the different ways to
// implement channels. Using go routines, or buffers to unblock the channel
// call and more.
func main() {
	// When creating a channel, you can unblock the channel after making it
	// by using a go routine to put the value on a channel. The print line
	// method takes off the channel.
	var c = make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	// Creating a buffer will allow the program to hold channels to be
	// retrieved from. Try and stay away from buffer channels.
	var c2 = make(chan int, 1)

	c2 <- 89
	fmt.Println(<-c2)
}
