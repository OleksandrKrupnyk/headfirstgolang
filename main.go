package main

import (
	"fmt"
	"magazine"
)

func printInfo(s *magazine.Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly rate:", s.Rate)
	fmt.Println("Active?:", s.Active)
}

func defaultSubscriber(name string) *magazine.Subscriber {
	var s magazine.Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}

func applyDiscount(s *magazine.Subscriber) {
	s.Rate = 4.99
}

func main() {

	subscriber1 := defaultSubscriber("Aman Sing")
	applyDiscount(subscriber1)
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Betn Ryan")
	printInfo(subscriber2)

	employee := magazine.Employee{Name: "Jorg", Salary: 500.0}
	address := magazine.Address{Street: "North Ave", City: "New York", State: "NY", PostalCode: "644888"}
	employee.Address = address
	fmt.Println(employee.City)
}
