package main

import "fmt"

func main() {
	// Here we will play a little with fmt package.
	i := 221
	fmt.Printf("%v\n", i) // Value of I
	fmt.Printf("%b\n", i) // Binary of i
	fmt.Printf("%x\n", i) // Hex of i
	fmt.Printf("%q\n", i) // Single quoted character
	fmt.Printf("%U\n", i) // Unicode Format

	s := fmt.Sprintf("%#x\t%b\t%x\t", i, i, i)
	fmt.Println(s)
}
