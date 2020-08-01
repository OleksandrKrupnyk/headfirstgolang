package main

import (
	"calendar"
	"fmt"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(1985)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(8)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
	fmt.Printf("Year is: %d\n", date.Year())
	fmt.Printf("Month is: %d\n", date.Month())
	fmt.Printf("Day is: %d\n", date.Day())
}
