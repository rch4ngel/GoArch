package main

import "fmt"

// Here we will practice slicing a slice. This practice of slicing
// is how we delete from a slice.
func main() {
	s := []int{0, 5, 3, 2, 5, 1}
	ss := s[0:3]

	fmt.Println(ss)
}
