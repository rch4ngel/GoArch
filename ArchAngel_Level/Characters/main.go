package Characters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type character struct {
	ID        string
	Firstname string
	Lastname  string
	Nickname  string
	Class     string
}

var cs = []character{}

func init() {
	LoadCharacters()
}

func main() {
	DisplayCharacters()
	SaveCharacters()
	ChannelCharacters()
}

func CreateCharacter(f string, l string, n string, c string) {
	var nc = character{"4", f, l, n, c}
	cs = append(cs, nc)
}

func LoadCharacters() {
	d, err := ioutil.ReadFile("Characters.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &cs)
	if err != nil {
		fmt.Println(err)
	}
}

func DisplayCharacters() {
	for _, c := range cs {
		fmt.Println(c.Firstname)
	}
}

func SaveCharacters() {
	f, err := os.Create("Characters.json")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	cj, err := json.Marshal(&cs)
	if err != nil {
		panic(err)
	}

	f.Write([]byte(cj))
}

func DeleteCharacters() {
	err := os.Remove("Characters.json")
	if err != nil {
		panic(err)
	}
}

func ChannelCharacters() {
	cc := make(chan []character)
	go sendCharacters(cc)
	receiveCharacters(cc)
}

func sendCharacters(c chan<- []character) {
	c <- cs
}

func receiveCharacters(c <- chan []character) {
	fmt.Println(<- c)
}