package main

import "fmt"

func main() {
	s := []int{3, 2, 1, 4, 5, 2, 4}

	// Technique for delete from a slice.
	s = append(s[:2], s[4:]...)

	fmt.Println(s)
}
