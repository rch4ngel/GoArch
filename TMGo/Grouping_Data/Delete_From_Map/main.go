package main

import "fmt"

// Here we will delete from a map. We will use the build in function
// delete.

func main() {
	m := map[string]int{
		"Jubei": 1,
		"Ryu":   2,
		"Clark": 29,
		"Lex":   32,
	}

	fmt.Println("And our contenders are! ")
	for k, v := range m {
		fmt.Println("Character:", k, "Rating:", v)
	}

	// Clark Kent knocks Lex Luther out.
	delete(m, "Lex")
	fmt.Println("Updating scoreboard! Jubei Knocks out Lex!")
	for k, v := range m {
		fmt.Println("Character:", k, "Rating:", v)
	}

	// Interesting note here. Deleting from the map reorders the map.
	// Perhaps this is due to a map being an un-ordered List.
}
