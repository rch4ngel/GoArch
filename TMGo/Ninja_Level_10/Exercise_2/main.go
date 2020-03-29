package main

import "fmt"

func main() {
	// c := make(chan<- int)
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
	fmt.Printf("c\t%T\n", c)
}
