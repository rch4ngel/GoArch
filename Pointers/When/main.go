package main

import "fmt"

// When to use pointers?
// User pointers when you have a huge chunk of data, and you don't want to pass
// that data around in your program. Instead pass the address around which is
// much smaller then passing around a large chunk of data.

// Another reason to use pointers is when the need to change a value at a certain location arises.

type beer struct {
	Name string
	Type string
	Abv  string
}

type beerCellar []beer

var bc beerCellar

func stockBeer() {
	for i := 0; i <= 100; i++ {
		b := beer{
			Name: fmt.Sprintf("Beer-%d", i),
			Type: fmt.Sprintf("Type-%d", i),
			Abv:  "11",
		}

		bc = append(bc, b)
	}
}

func beerInventory() *beerCellar {
	return &bc
}

func beerLocation(b *beer) {
	fmt.Printf("%v  \t%d\t%v\t%b\n", b.Name, &b, &b, &b)
}

func allBeerLocations(sb ...beer) {
	for _, b := range sb {
		beerLocation(&b)
	}
}

func init() {
	stockBeer()
}

func main() {
	bi := beerInventory()
	fmt.Println("Location of the beer brewery")
	fmt.Println(&bi)
	fmt.Println("Locations of all the beer in the inventory")
	allBeerLocations(bc...)
}