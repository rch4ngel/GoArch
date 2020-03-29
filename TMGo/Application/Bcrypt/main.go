package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// Never store raw passwords... Always store passwords encrypted. Bcrypt is used to do this.
// Bcrypt is not yet checked into the standard library but it is in the x directory.

func main() {
	pw := `MyPassword`
	bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pw)
	fmt.Println(string(bs))
}
