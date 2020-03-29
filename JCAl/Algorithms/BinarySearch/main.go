package main

import (
	"bytes"
	"fmt"
)

func main() {
	l := []int{43, 22, 1, 46, 7, 5, 4, 3, 334, 6}
	fmt.Println("Unsorted List:", l)
	bubbleSort(l)
	fmt.Println("Sorted List:", l)
	binarySearch(l, 3)

	i := binarySearch(l, 4)
	if i == -1 {
		fmt.Println("Item not found!")
	} else {
		fmt.Println("Found the item:", l[i])
	}
}

func binarySearch(sl []int, theItem int) int {
	lw := 0
	hi := len(sl) - 1

	for lw <= hi {
		m := lw + (hi-lw)/2
		mv := sl[m]
		fmt.Println("Half the number of items:", m)
		fmt.Println("Looking for", theItem)
		fmt.Println("The narrowing list:", mv)
		if mv == theItem {
			return m
		} else if mv > theItem {
			hi = m - 1
		} else {
			lw = m + 1
		}
	}

	return -1
}

func bubbleSort(ul []int) {
	for i := 0; i < len(ul); i++ {
		if !swap(ul, i) {
			return
		}
	}
}

func swap(ul []int, pp int) bool {
	sz := len(ul) - pp
	fi := 0
	si := 1
	ds := false

	for si < sz {
		fv := ul[fi]
		sv := ul[si]

		if isGreaterThan(fv, sv) {
			ds = true
			ul[fi] = sv
			ul[si] = fv
		}

		fi++
		si++
	}

	return ds
}

// isGreaterThan is a helper function that will be type agnostic, well for strings and integers.
func isGreaterThan(ft interface{}, st interface{}) bool {
	switch ft.(type) {
	case string:
		fs := fmt.Sprintf("%s", ft)
		ss := fmt.Sprintf("%s", st)

		fxb := []byte(fs)
		sxb := []byte(ss)

		if bytes.Compare(fxb, sxb) == 1 {
			fmt.Println(fxb, "is greater than", sxb)
			return true
		}
	case int:
		fs := fmt.Sprintf("%d", ft)
		ss := fmt.Sprintf("%d", st)

		fxb := []byte(fs)
		sxb := []byte(ss)

		if bytes.Compare(fxb, sxb) == 1 {
			fmt.Println(fxb, "is greater than", sxb)
			return true
		}
	default:
		return true
	}

	return false
}
