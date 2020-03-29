package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Let it be known that when I started this file I came up with the famous go
// application GoTime or ItsGoTime.... 10/7/18

// In this example we are creating a race condition. The way this works, is:
// Multiple Functions or Methods are referencing a single value at the same time.
// Below multiple go routines are accessing counter variable. The result shows,
// that the program is returning the values in an unordered fashion. This is because
// each go routine is returning the value at different times.

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	counter := 0
	for i := 0; i < gs; i++ {
		// Setting up a go routine with a function literal.
		go func() {
			v := counter
			v++
			counter = v
			// At the end of the function, setting wg.Done() to
			// Show that goroutine is done.
			fmt.Println(counter)
			wg.Done()
		}()
	}
	wg.Wait()
}
