package main

import "fmt"

// Anonymous structs work by not having a name.
// declare anonymous structs within a function.
func main() {
	// Anonymous struct
	hl := struct {
		Nickname      string
		AngelAblility string
	}{
		Nickname:      "Jubie",
		AngelAblility: "Sword of Benevolence",
	}

	fmt.Println(hl)
}
