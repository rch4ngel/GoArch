package main

import "fmt"

// Methods accept a receiver in the function statement. So methods are attached
// to a type.

// func (r receiver) identifier(parameters) (return(s)) {code}

type person struct {
	Firstname string
	Lastname  string
}
type archAngel struct {
	person
	AngelAbility string
}

func (a archAngel) speak() {
	fmt.Println("Hi! I am", a.Firstname, a.Lastname, "My angel ability is", a.AngelAbility)
}

func main() {
	aa := archAngel{
		person: person{
			Firstname: "Jubei",
			Lastname:  "Kibagami",
		},
		AngelAbility: "Sword of Benevolence",
	}

	fmt.Println(aa)
	aa.speak()
}
