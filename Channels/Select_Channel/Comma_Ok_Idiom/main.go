package main

import "fmt"

// Here we will explore the comma ok idiom.

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// This will fire off a go routine to send the defined arguments
	go send(eve, odd, quit)
	// This could will block until all values are received. No need for a wait group or a buffer channel.
	receive(eve, odd, quit)
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the Even Channel:", v)
		case v := <-o:
			fmt.Println("From the Odd Channel:", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("Comma Ok Idiom reports its not Ok... RETURN!", i, ok)
				return
			}else {
				fmt.Println("Comma Ok Idiom reports its Ok", i, ok)
			}
		}
	}
}
