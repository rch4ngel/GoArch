package main

// This is a more real world example of closure...

func incrementer() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

