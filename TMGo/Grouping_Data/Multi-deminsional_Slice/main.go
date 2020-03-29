package main

import "fmt"

func main() {
	s1 := []string{"Hello", "First", "Slice", "Of", "Strings"}
	s2 := []string{"I", "Am", "The", "Second", "Slice", "Of", "Strings"}

	ms := [][]string{s1, s2}
	fmt.Println(ms)
}
