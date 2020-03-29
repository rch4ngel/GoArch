package main

import (
	"fmt"
	"runtime"
	"sync"
)

type point struct {
	ID        string
	Northing  float64
	Easting   float64
	Elevation float64
}

type line struct {
	point
	Bearing float64
	Length  float64
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	p := PointCreator()
	fmt.Println(runtime.NumGoroutine())
	go LineCreator(p)
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
}

func PointCreator() *point {
	p := point{
		ID:        "1",
		Northing:  3215.4,
		Easting:   44632.1,
		Elevation: 432.2,
	}
	fmt.Println(p)

	return &p
}

func LineCreator(p *point) {
	l := line{
		point:   *p,
		Bearing: 22,
		Length:  1000,
	}
	fmt.Println(l)
	wg.Done()
}
