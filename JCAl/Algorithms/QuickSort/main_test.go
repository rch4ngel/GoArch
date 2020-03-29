package main

import (
	"fmt"
	"testing"
)

func TestLoadSensors(t *testing.T) {
	sm := LoadSensors()
	for s := range sm {
		fmt.Println("Equipment:", s)
		//
		//for j := 0; j < len(sm["DR02"]); j++ {
		//	fmt.Println("\t\tSensor", sm["DR02"][j])
		//}
	}
}
