package main

import (
	"calendar"
	"fmt"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetDay(3)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(8)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2020)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetTitle("Hello me!")
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
	fmt.Println(event.Title())
}
