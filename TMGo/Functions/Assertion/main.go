package main

import "fmt"

type original int

type person struct {
	Firstname string
	Lastname  string
}

type hero struct {
	Name string
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("Type person speaking here.")
}

func (h hero) speak() {
	fmt.Println("Type hero speaking here.")
}

func main() {
	p := person{
		Firstname: "Jubei",
		Lastname:  "Kibagami",
	}

	h := hero{
		"Vagabond",
	}

	i := 42

	SwitchOnType(p)
	SwitchOnType(h)
	SwitchOnType(i)
}

// Below will show the usage of assertion.
// when evaluating x below, you cannot
// call x.Name from the start. Assertion is
// required to expose the type's properties.
func SwitchOnType(x interface{}) {
	switch x.(type) {
	case hero:
		fmt.Println("Type is hero")
		fmt.Printf("%s\n", x.(hero).Name)
	case person:
		fmt.Println("Type is person")
		fmt.Printf("%s, %s\n", x.(person).Firstname, x.(person).Lastname)
	case original:
		fmt.Println("Type is original")
	default:
		fmt.Println("Default")
	}
}
