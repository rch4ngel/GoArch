package main

import "fmt"

// Select statements is another tool for using channels.
func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)
	receive(eve, odd, quit)

	fmt.Println("About to exit.")
}

func send(e, o, q chan<- int) {
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From The Eve Channel:", v)
		case v := <-o:
			fmt.Println("From The Odd Channel:", v)
		case v := <-q:
			fmt.Println("From The Quit Channel:", v)
			return
		}
	}
}
