package main

import "fmt"

// Here are the ways to declare a string.
var s1 string = `Hello "Are you there" No?`
var s2 string = "Hi"
var s4 string // This bad boy will be assigned to zero type

func main() {
	s3 := "Greeting from inside the main function."
	// Here are some different fmt.print functions.
	fmt.Println(s1)
	// Printf %T will print out the type! Bad ass!
	fmt.Printf("%T\n", s2)
	fmt.Println(s3)
	fmt.Println(s4)

	// Strings in go are immutable which means that when it
	// is created it cannot be changed. It can be assigned
	// another value but the bytes representing the string cannot
	// be changed.

	// Strings are encoded using UTF-8 encoding. A utf-8 code point
	// is 1-4 bytes. Each code point coresponds with a character.
	// uint8 is actually an alias for byte. 8 bits represent and byte.
	//

	for i, bs := range s3 {
		fmt.Printf("Index %d, has binary %b which is %x in hexadecimal\n", i, bs, bs)
		fmt.Printf("UTF-Encoding for byte %d is %#U\n", bs, bs)
	}
}
