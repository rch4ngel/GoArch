package main

import "fmt"

// Channels can be assigned as well. However it can only be done from general
// To more specific. See below

func main() {
	cs := make(chan<- int, 1)
	cr := make(<-chan int, 1)
	c := make(chan int, 1)

	//c = cs this is taking a specific and assigning it to a general channel NOT ALLOWED!
	//c = cr same thing here, cr is specific and we are trying to assign it to a general channel
	cr = c // Works! c is general assigning it to a more specific value.
	cs = c
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
}
