package main

import (
	"fmt"
	"liquids"
)

type Number int

func (n Number) Double() {
	n *= 2
}

func (n *Number) Display() {
	fmt.Println(*n)
}

func main() {
	krazFuel := liquids.Liters(501.234)
	fmt.Printf("KrAZ fuel is: %0.3f\n", krazFuel)
	fordFuel := krazFuel.ToGallons()
	fmt.Printf("Ford fuel is: %0.3f\n", fordFuel)
	number := Number(4)
	number.Double()
	number.Display()
}
