package main

import "fmt"

// The Original Bubble Sort implemented prior has some issues in performance.
// Other than it being BigO n^2, the issues in the original implementation is it
// will continue to iterated over sorted values. This implementation will not.

func main() {
	xn := []int{5, 5, 4, 3, 2, 1}
	BubbleSort(xn)
	fmt.Println("The Final List of Numbers")
}

func BubbleSort(xi []int) {
	fmt.Println("The Original List of Numbers:", xi)
	for i := 0; i < len(xi); i++ {
		if !Swap(xi, i) {
			return
		}
		fmt.Println("List After The First Swap:", xi)
	}
}

// Swap up until the previous pass. This addition is easy to implement
// just a simple parameter to pass in. The bool return will notify
// the caller that it has or hasn't swapped a number.
func Swap(xi []int, pp int) bool {
	n := len(xi)
	fi := 0
	si := 1
	ds := false

	for si < n-pp {
		fn := xi[fi]
		sn := xi[si]

		if fn > sn {
			xi[fi] = sn
			xi[si] = fn
			ds = true
		}

		fi++
		si++
	}

	return ds
}
