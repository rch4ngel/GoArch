package main

import "fmt"

func main() {
	x := 421
	if x == 42 {
		fmt.Println("Our value is 42")
	} else if x == 420 {
		fmt.Println("Our value is 420")
	} else if x == 69 {
		fmt.Println("Our value is 69")
	} else {
		fmt.Println("Our value is not 42, 420 or 69")
	}

}