package main

import "fmt"

type Character struct {
	Name    string
	Ability string
}

type Characters []*Character

func main() {
	xc := PopulateCharacters()
	xc.DisplayCharacters()
}

func PopulateCharacters() *Characters{
	c := make(chan *Character)
	go sendCharacters(c)
	rc := receiveCharacters(c)

	return rc
}

func (xc *Characters) DisplayCharacters() {
	for _, c := range *xc {
		fmt.Printf("Character: %s\n", *c)
	}
}

func sendCharacters(c chan<- *Character) {
	const ncs = 100
	for i := ncs; i > 0; i-- {
		name := fmt.Sprintf("Character-%d", i)
		ability := fmt.Sprintf("Ability-%d", i)
		c <- &Character{name, ability}
	}
	close(c)
}

func receiveCharacters(c <-chan *Character) *Characters {
	xc := Characters{}
	for v := range c {
		xc = append(xc, v)
	}
	return &xc
}
