package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	length, width := 1.0, 2.5
	height := 25
	var now time.Time = time.Now()
	fmt.Println(
		"Hello",
		strings.Title("awesome string"),
		reflect.TypeOf(31.9),
		length+width*float64(height),
		now.Day())

}
