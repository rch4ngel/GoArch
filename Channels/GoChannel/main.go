package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"io/ioutil"
	)

type Sensor struct {
	Name  string
	Value float64
}

var wg sync.WaitGroup

func LoadSensors() {
	const ns = 320
	var xs []Sensor
	nf, err := os.Create("Sensors.json")
	if err != nil {
		fmt.Println(err)
	}

	defer nf.Close()

	for i := 0; i < ns; i++ {
		rn := rand.ExpFloat64()

		s := Sensor{
			Name:  fmt.Sprintf("Sensor-%d", i),
			Value: rn,
		}

		xs = append(xs, s)
	}

	sj, err := json.Marshal(xs)
	if err != nil {
		fmt.Println("Error Parsing Go Object to JSON\n", err)
	}

	nf.Write(sj)

	wg.Done()
}

// This functionality still needs to be defined.
// Was not completed due to wanting to learn directional channels.
func FindSensor(s string) {
	var xs []Sensor
	sf, err := ioutil.ReadFile("Sensors.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(sf, &xs)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	wg.Add(1)
	go LoadSensors()
	wg.Wait()

	FindSensor("Sensor-1")
}

