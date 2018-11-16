package main

import "fmt"

func main() {
	if true {
		fmt.Println("This is true statement 1")
	}

	if false {
		fmt.Println("This is false statement 1")
	}

	if !true {
		fmt.Println("This is not true (!true) statement 1")
	}

	if !false {
		fmt.Println("This is not false (!false) statement 1")
	}
}

