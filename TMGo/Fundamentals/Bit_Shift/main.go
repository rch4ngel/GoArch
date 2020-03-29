package main

import "fmt"

// Bit shifting is the process of shifting a bit 0 or 1 from either left
// or right.

// iota is a great tool used for bit shifting. below is an example of
// how to represent kb, mb, and gb using bit shifting with iota
const (
	_ = iota             // Throw the zero value away..
	k = 1 << (iota * 10) // from mb -> kb -> gb 10 zeros are added.
	m = 1 << (iota * 10)
	g = 1 << (iota * 10)
)

func main() {
	i := 2
	fmt.Printf("%d, %b\n", i, i)

	j := 1 << 2
	fmt.Printf("%d, %b\n", j, j)

	// How about kb, mb, gb ?
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("A kilobyte %d in binary %b, and in hexadecimal %x\n", kb, kb, kb)
	fmt.Printf("A megabyte %d in binary %b, and in hexadecimal %x\n", mb, mb, mb)
	fmt.Printf("A gigabyte %d in binary %b, and in hexadecimal %x\n", gb, gb, gb)

	// Testing out the bit shifting technique written above
	fmt.Println("Testing out the handy dandy bit shifter!")
	fmt.Printf("A kilobyte %d in binary %b, and in hexadecimal %x\n", k, k, k)
	fmt.Printf("A megabyte %d in binary %b, and in hexadecimal %x\n", m, m, m)
	fmt.Printf("A gigabyte %d in binary %b, and in hexadecimal %x\n", g, g, g)
}
