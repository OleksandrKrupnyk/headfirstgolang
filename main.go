package main

import (
	"fmt"
)

type Liters float64
type Gallons float64

func main() {
	carFuel := Gallons(Liters(240) * 0.264)
	busFuel := Liters(Gallons(10.0) * 3.785)
	fmt.Printf("Gallons %0.1f Liters %0.1f", carFuel, busFuel)
}
