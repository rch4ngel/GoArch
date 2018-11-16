package main

import "fmt"

// Working with embedded types. Don't name the structure that is embedded, just provide its type. When declaring and assigning
// the embedded type, you do need to provide the field name.
// Another item for discussion will be instances where a name collision occurs (When same name appears in both structs where
// the embedding occurs). You can access it using dot notation. <outer-type>.<inner-type>.firstname for example.

type person struct {
	Firstname string
	Lastname  string
	Age       int
	Gender    string
}

type halCharacter struct {
	person
	Nickname     string
	Class        string
	AngelAbility string
	PowerLevel   int
}

func main() {
	cl1 := halCharacter{
		person: person{
			Firstname: "Jubei",
			Lastname:  "Kibagami",
			Age:       33,
			Gender:    "m",
		},
		Nickname:     "Jubei",
		Class:        "Samurai",
		AngelAbility: "Sword of Benevolence",
		PowerLevel:   1000,
	}

	fmt.Println(cl1)

}
