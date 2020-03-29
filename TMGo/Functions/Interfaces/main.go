package main

import "fmt"

// A value can be more of one type. A value can be more of one type
// a value can be more of one type.

// Polymorphism! Interfaces allow us to define behavior.
// So let us get to it.

type person struct {
	Firstname string
	Lastname  string
}

type archAngel struct {
	person
	AngelAbility string
}

// So what this means...
// Any type that uses the speak function
// will be of type human interface.
type human interface {
	speak()
}

// Is of type human interface
func (a archAngel) speak() {
	fmt.Println("Hi I am", a.Firstname, a.Lastname)
}

func (p person) speak() {
	fmt.Println("Person is now of type human interface... Human is", p.Firstname, p.Lastname)
}

func describe(h human) {
	fmt.Println("Any type of type human can be used here!")
}

func main() {
	aa := archAngel{
		person: person{
			Firstname: "Ryu",
			Lastname:  "",
		},
		AngelAbility: "Hadouken",
	}

	aa.speak()

	p := person{
		Firstname: "Jubei",
		Lastname:  "Kibagami",
	}

	p.speak()

	// Since type archangel and person implement the human interface
	// so we can pass in archangel type, and person type as arguments.
	describe(aa)
	describe(p)
}
