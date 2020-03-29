package main

import "fmt"

// To keep thses files a little concise here is just an example of a receive channel.

func main() {
	cr := make(<- chan int, 1)
	fmt.Printf("%T\n", cr)
}
