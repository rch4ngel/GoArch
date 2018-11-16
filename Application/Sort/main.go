package main

import (
	"fmt"
	"sort"
)

// Moving onto the next stage dipping into the use of standard library interfaces.
// It will be wise to study the implementations here to see how the best of the best
// Optimize systems.

func main() {
	xi := []int{4, 1, 5, 2, 5, 6, 6, 3, 23, 542, 234, 3, 5}
	sort.Ints(xi)

	fmt.Println(xi)

	xs := []string{"Jason", "Carl", "Mike", "John"}
	sort.Strings(xs)
	fmt.Println(xs)
}
