package main

import "fmt"

// Since functions are first class functions you can
// pass them in as callbacks. This is also called functional programming.
// functional programming is not recommended in go.

func main() {
	ii := []int{1, 3, 4, 5, 2, 5}

	s := sum(ii...)

	fmt.Println(s)
	fmt.Println(even(sum, ii...))
	fmt.Println(odd(sum, ii...))
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}

	return total
}

// Func even that uses functional programming pattern.
func even(f func(xi ...int) int, vi ...int) int {
	var si []int
	for _, v := range vi {
		if v%2 == 0 {
			si = append(si, v)
		}
	}

	return f(si...)
}

// Func odd that uses functional programming pattern.
func odd(f func(xi ...int) int, vi ...int) int {
	var si []int
	for _, v := range vi {
		if v%2 != 0 {
			si = append(si, v)
		}
	}

	return f(si...)
}
