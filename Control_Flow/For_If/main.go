package main

import "fmt"

func main() {
	fmt.Println("Welcome!")
	// Here we will use a for loop that will print out all even numbers up to 100
	i := 0
	for {
		i++
		if i > 100 {
			break
		}
		// If it is odd, it will continue back up to the beginning and not
		// run any of the code below.
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}

