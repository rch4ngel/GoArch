package main

import "fmt"

// Bubble sort implementation. The bubble sort will have 2 functions, because practically it has 2 functions.
// Functions are:
//					Loop
//					Sweep
// Loop will iterate the size of the slice n
// Sweep will iterate the size n - 1

func main() {
	xn := []int{3, 2, 2, 4, 7, 3, 6, 9, 0}
	bubbleSort(xn)
}

// func BubbleSort will represent the loop function of our program.
func bubbleSort(xn []int) {
	fmt.Println("Slice of numbers to be sorted:", xn)

	for i := 0; i < len(xn); i++ {
		sweep(xn)
		fmt.Println("Slice of numbers after the first pass:", xn)

	}
	fmt.Println("End Bubble Sort Result:", xn)
}

func sweep(xi []int) {
	n := len(xi)
	fi := 0
	si := 1

	// The way to write a while loop in go.
	// here the loop will loop through while
	// the second index is less than size of slice.
	for si < n {
		fn := xi[fi]
		sn := xi[si]
		fmt.Println("Comparing numbers:", fn, sn)

		if fn > sn {
			xi[fi] = sn
			xi[si] = fn
		}
		fi++
		si++
	}
}
