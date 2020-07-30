package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {

	var subscriber1 subscriber
	subscriber1.name = "Aman Sing"
	fmt.Println("Name:", subscriber1.name)
	var subscriber2 subscriber
	subscriber2.name = "Betn Ryan"
	fmt.Println("Name:", subscriber2.name)

}
