package main

import (
	"fmt"
	"sort"
)

// Custom sort examples and implementations. The trick here is to implement the sort interface, interface Interface.
// Len, Less, Swap.
// After implementing the interface, it can be used on numerous properties.

type person struct {
	Firstname string
	Lastname  string
	Age       int
}

type byAge []person
type byName []person

func (ba byAge) Len() int {
	return len(ba)
}

func (ba byAge) Less(i int, j int) bool {
	return ba[i].Age < ba[j].Age
}

func (ba byAge) Swap(i int, j int) {
	ba[i], ba[j] = ba[j], ba[i]
}

func (bn byName) Len() int {
	return len(bn)
}

func (bn byName) Less(i int, j int) bool {
	return bn[i].Lastname < bn[j].Lastname
}

func (bn byName) Swap(i int, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}

func main() {
	p1 := person{Firstname: "Clint", Lastname: "Bar", Age: 33}
	p2 := person{Firstname: "Lil", Lastname: "Brandt", Age: 22}
	p3 := person{Firstname: "Cindy", Lastname: "Grant", Age: 40}

	people := []person{p1, p2, p3}
	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)

	sort.Sort(byName(people))
	fmt.Println(people)
}
