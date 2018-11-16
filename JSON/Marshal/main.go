package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Firstname string
	Lastname  string
	Gender    string
	Age       int
}

var people []person

func main() {
	p1 := person{
		Firstname: "Jubie",
		Lastname:  "Kibagami",
		Gender:    "Male",
		Age:       32,
	}
	p2 := person{
		Firstname: "Ryu",
		Lastname:  "",
		Gender:    "Male",
		Age:       31,
	}
	p3 := person{
		Firstname: "Luke",
		Lastname:  "Skywalker",
		Gender:    "Male",
		Age:       26,
	}

	people = append(people, p1, p2, p3)
	peopleJSON()
}

func peopleJSON() {
	pj, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(pj))
}
