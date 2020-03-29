package main

import "fmt"

// Slices are of variable length and they have an underlying type of Array. Slices have a:
//
// Pointer - Will Point to the first reachable value in the array. Not always the first value.
// Length - Will Report the Length of the array, will never be larger than the capacity.
// Capacity - Will Report the capacity of the underlying array.

func main() {
	// Setting the size will set the type to an array
	a := [2]int{}
	fmt.Println(len(a))
	fmt.Println(cap(a))

	xa := make([]int, 5, 6)
	xa[0] = 22
	xa[1] = 33
	xa[2] = 44
	xa[3] = 55
	xa[4] = 66
	// xa[5] = 77 Out of Range! We fix this by using append...
	xa = append(xa, 77, 88) // This will dynamically resize the underlying array size or capacity to double
	  							  // the size of the original. So it will have a capacity of 12.

    // xa[7] = 99 This will still panic with an out of range error.... This is because the size will only be the number of elements in the SLICE..
    fmt.Println("Size of our slice of ints:", len(xa))
	fmt.Println("Capacity of our slice of ints:", cap(xa))
	fmt.Println(xa)
}


