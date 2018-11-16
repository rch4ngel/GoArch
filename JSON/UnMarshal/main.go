package main

import (
	"encoding/json"
	"fmt"
)

// Unmarshalling JSON. To unmarshal JSON there needs to be a []bytes, and a value pointed to by an interface v.

type person struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Gender    string `json:"Gender"`
	Age       int    `json:"Age"`
}



func packPeople() []byte {
	var p person
	var people []person

	for i := 0; i <= 20; i++ {
		p = person{
			Firstname: fmt.Sprintf("Clone-%d", i),
			Lastname:  fmt.Sprintf("Kent"),
			Gender:    "Male",
			Age:       22,
		}
		people = append(people, p)
	}

	pj, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	return pj
}

func unpackPeople(pj []byte) []person{
	var people []person
	err := json.Unmarshal(pj, &people)
	if err != nil {
		fmt.Println(err)
	}
	return people
}

func main() {
	pp := packPeople()
	fmt.Println(string(pp))
	up := unpackPeople(pp)
	fmt.Println(up)
}
