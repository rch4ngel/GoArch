package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Welcome to the goground where concurrent magic will ensue....

func main() {
	Curator()
}

// Creating a waitgroup of size 100 to
// "hold" the number of go routines.
const gs = 100

// When setting mutex for forbidding race conditions or limiting them
// Be sure to think how the rest of the program makes use of it, even a
// print line statement.
func Curator() {
	fmt.Println("Curator is fetching your data.")
	GetMutexCounter()
	GetAtomicCounter()
}

func GetMutexCounter() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gs)
	// Looping through all go routines,
	// spinning off 100 different go routines.
	counter := 0
	for i := 0; i < gs; i++ {
		go func() {
			// Mutex lock the race condition. There are numerous ways
			// To do this, which will be shown below.
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter) // This has to placed within the lock as well since it is
			mu.Unlock()          // accessing the counter variable as well.
			wg.Done()
		}()
	}
	// Wait until the go routines are all done.
	wg.Wait()
}

func GetAtomicCounter() {
	var wg sync.WaitGroup

	wg.Add(gs)
	var ac int64
	for i := 0; i < gs; i++ {
		// Atomic has a bit more magic with it. It needs a pointer to an int and a change in
		// state. In this case we are adding the int64 by 1.
		atomic.AddInt64(&ac, 1)
		runtime.Gosched()
		fmt.Println(ac)
		wg.Done()
	}
	wg.Wait()
}
