package main

import (
	"datafile"
	"fmt"
	"log"
)

func main() {

	numbers, err := datafile.GetFloats("C:\\Users\\sasha\\go\\src\\awesomeProject\\data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(numbers)
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)

}
