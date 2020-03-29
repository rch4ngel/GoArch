package main

import "fmt"

// Working on the structures here. Let us begin...
// Hal == "Hero Afterlife"
type HalCharacters struct {
	Name         string
	Class        string
	AngelAbility string
	PowerLevel   int
}

func main() {
	hl := HalCharacters{
		Name:         "Jubei",
		Class:        "Samurai",
		AngelAbility: "Sword of Benevolence",
		PowerLevel:   1000,
	}

	fmt.Println(hl)

	// Dragonball type Reading Level info
	fmt.Println("Name:", hl.Name, "\nClass:", hl.Class, "\nAngelAbility:", hl.AngelAbility, "\nPowerLevel:", hl.PowerLevel)
}
