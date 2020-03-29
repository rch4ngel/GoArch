package main

import (
	"fmt"
	"math"
)

// Method sets are the methods attached to a given type.
// Non pointer receivers work with values that are POINTERS or NON-Pointers
// POINTER receivers only work with values that are POINTERS

type circle struct {
	Radius float64
}

type shape interface {
	area() float64
}
// Non-Pointer Receiver
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c *circle) circleArea() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func circleInfo(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{
		Radius: 5,
	}

	info(c)
	circleInfo(&c)
}
