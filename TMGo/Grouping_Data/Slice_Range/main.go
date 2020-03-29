package main

import "fmt"

func main() {
	s := []int{1, 3, 5, 4, 2, 4, 2, 1, 10}
	for i, v := range s {
		fmt.Println("Index: ", i, " Has a value of: ", v)
	}
}
