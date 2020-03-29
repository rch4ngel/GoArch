package main

import (
	"fmt"
	"sync"
)

// Fan in is a design pattern for channels. Fanning in is where multiple channels are
// merged into one channel or fanned into one channel. In this example even and odd channels are fanned into the
// fanin channel.

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)
	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}
}

func send(e, o chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func receive(e, o <-chan int, f chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	// Both these go routines will be sent off at the same time to compute the programmed instructions.
	// This code will block here until the code is done ranging.
	go func() {
		for v := range e {
			f <- v
		}
		wg.Done()
	}()
	// As above this will block until the range is closed.
	go func() {
		for v := range o {
			f <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(f)

}
