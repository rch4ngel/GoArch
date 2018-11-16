package main

import "fmt"

// iota is a pre declared identifier.
// here iota will increment by one. the
// next constant will restart the counter.
const (
	a = iota
	b = iota
	c = iota
)

const (
	d = iota
	e = iota
	f = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T", a)
	fmt.Printf("%T", b)
	fmt.Printf("%T", c)
}
