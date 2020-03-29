package main

import "fmt"

type character struct {
	Name         string
	AngelAbility string
}

var xc []character

func main() {
	GenerateCharacters()
	DisplayCharacters()
}

func GenerateCharacters() {
	sc := make(chan character)
	go sendCharacters(sc)
	receiveCharacters(sc)
}

func sendCharacters(c chan<- character) {
	const nc = 100
	for i := 0; i < nc; i++ {
		name := fmt.Sprintf("Character-%d", i)
		ability := fmt.Sprintf("Ability-%d", i)
		c <- character{Name: name, AngelAbility: ability,}
	}
	close(c)
}

func receiveCharacters(c <-chan character) {
	for v := range c {
		xc = append(xc, v)
	}
}

func DisplayCharacters() {
	for _,v := range xc {
		fmt.Println("Character:", v.Name, " \tAngelAbility:", v.AngelAbility)
	}
}
