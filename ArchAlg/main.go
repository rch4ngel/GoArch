package main

import (
	"fmt"
	"github.com/archangel/ArchAlg/DataModels"
)

func main() {
	xu := user.GetUsers()
	for _, u:= range xu {
		fmt.Println(*u)
	}
}
