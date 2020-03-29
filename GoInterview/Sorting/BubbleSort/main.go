package main

import (
	"fmt"
	"os"
	)

// Bubble Sort Implementation. The Bubble Sort is not an effective sort and runs in O(log(n^2)) time.
// How the Bubble sort works is that it compares a pair at a time and swaps if the first value is larger
// than the second value. After One pass the bubble sort will have put the largest value at the end of the array.

func main() {
	xi := []int{55, 88, 2, 33, 24, 100, 442, 1}
	fmt.Println(xi)
	bubbleSort(xi)
	fmt.Println(xi)
	writeToFile(xi)
}

func sweep(numbers []int, previousPass int) {
	N := len(numbers)
	fi := 0
	si := 1

	for si < N-previousPass {
		fn := numbers[fi]
		sn := numbers[si]
		fmt.Println("Comparing The Following:", fn, sn)

		if fn > sn {
			numbers[fi] = sn
			numbers[si] = fn
		}
		fi++
		si++
	}
}

func bubbleSort(numbers []int) {
	N := len(numbers)
	for i := 0; i < N; i++ {
		sweep(numbers, i)
	}
}

func writeToFile(sortedNumbers []int) {
	f, err := os.Create("sorted_numbers.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, i := range sortedNumbers {
		is := fmt.Sprintf("%d ", i)
		f.WriteString(is)
	}
}
