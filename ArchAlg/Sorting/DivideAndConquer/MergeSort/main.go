package main

import (
	"fmt"
	"math"
)

// Algorithm Type		:		Sorting
// Algorithm			: 		Merge Sort
// Formal Definition	:
//								Input	: 	A sequence of n numbers {a1, a2, ..... an}
//								Output	:	A permutation (reordering) {a1, a2, ...... an} of the input sequence each that
//								            a1 <= a2 <= ...... <= an

// Legend
// c represents a constant time
// O represents BigO Notation
// n represents the number of items of input

// Below is the pseudocode of the merge sort. The merge sort will run in O(n*lg(n))
// MERGE(A, p, q, r)
//  n1 = q - p + 1
//  n2 = r - q
//  let l1[1... n1 + 1] and l2[1... n2 + 1]
//  for i = 1  to n1 {
//    l1[i] = A[p + i -1]
//  }
//  for j = 1 to n2 {
//    li[j] = A[q + j]
// }

// Below is the textual representation of the process of the merge sort

// 1 Based index
//          [1    2   3   4   5   6   7   8]
// Original [55, 44, 33, 22, 22, 33, 44, 55]



func main() {
	xi := []int{55, 44, 33, 22, 22, 33, 44, 55}
	p := 0
	q := len(xi) / 2
	r := len(xi)

	MergeSort(xi, p, r)
	Merge(xi, p, q, r)
	fmt.Println(xi)
}

func Merge(xi []int, p, q, r int) {
	n1 := q - p
	n2 := r - q

	lxi := make([]int, n1 + 1)
	rxi := make([]int, n2 + 1)

	for i := 0; i < n1; i++ {
		lxi[i] = xi[p + i]
	}
	for j := 0; j < n2; j++ {
		rxi[j] = xi[q + j]
	}

	lxi[n1] = int(math.Pow(10, 10))
	rxi[n2] = int(math.Pow(10, 10))

	fmt.Println("Left Slice:", lxi)
	fmt.Println("Right Slice:", rxi)

	i := 0
	j := 0


	for k := p; k < r; k++ {
		if lxi[i] <= rxi[j] {
			xi[k] = lxi[i]
			i++
		} else {
			xi[k] = rxi[j]
			j++
		}
	}
	fmt.Println("Main Slice:", xi)
}

func MergeSort(xi []int, p, r int) {
	if p < r {
		q := (p + r) /2
		MergeSort(xi, p, q)
		MergeSort(xi, q + 1 , r)
		Merge(xi, p, q, r)
	}
}
