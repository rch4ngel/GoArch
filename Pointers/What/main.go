package main

import "fmt"

// Pointers "point" to a memory location. To see the address you can use &<pointer-name>
// The & operator will give the address where the value is stored.
// The * operator will give the value at an address.
// Get the address &<variable>
// Dereference address syntax *<variable>

func main() {
	i := 3
	fmt.Println(i)
	fmt.Println(&i)

	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", &i)

	p := &i // P is now assigned the same address as i...
	fmt.Println(p)
	fmt.Println(*p)
	// Assign the value at address p to 4...
	*p = 4
	// Does this change the value of i as well?
	fmt.Println(i)
	fmt.Println(*p)

}
