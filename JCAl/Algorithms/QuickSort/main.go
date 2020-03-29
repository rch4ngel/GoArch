package main

import (
	"fmt"
	"github.com/archangel/JCAl/Algorithms/QuickSort/sensors"
)

// Below is the implementation of the QuickSort. We could use the sort package,
// but lets see how much slower my implementation is.

func main() {
	xs, err := sensors.AllSensors()
	if err != nil {
		fmt.Println(err)
	}

	for _, s := range xs {
		fmt.Println(s)
	}
}


