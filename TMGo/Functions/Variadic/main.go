package main

// Variadic parameters take in zero or more arguments. When using
// variadic parameters, you can pass in not variadic parameters,
// however those parameters must be passed in first.

import "fmt"

func main() {
	sum(1, 2, 4, 5, 3, 2)
	avg("Jubei", 1, 4, 5, 6, 3, 4, 2)
}

func avg(s string, x ...int) {
	count := 0
	sum := 0
	for i, v := range x {
		count += 1
		sum += v
		fmt.Println("Index:", i, "has value:", v, "count is:", count, "sum is:", sum)
	}
	fmt.Println("You passed in string:", s)
	fmt.Println("Total is:", sum)
	fmt.Println("Average is:", sum/count)
}

func sum(x ...int) {
	s := 0
	for i, v := range x {
		fmt.Println("Index:", i, "With Value: ", v)
		s += v
	}

	fmt.Println("The sum is: ", s)
}
