package main

import "fmt"

// Range will range over the channels until the channel is closed.

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
