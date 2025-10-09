package main

import (
	"challenges"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Type Assertions!")

	var name interface{} = "Abcd"
	var age interface{} = 24

	dynamicDev := challenges.GetDeveloper(name, age)
	fmt.Println(dynamicDev.Name)
	fmt.Println(dynamicDev.Age)

	// Parse date of birth (format: YYYY-MM-DD)
	dob, err := time.Parse("2006-01-02", "2001-04-15")
	if err != nil {
		panic(err)
	}

	fmt.Println("getEngg!")
	getEngg(dynamicDev.Name, dob)

	fmt.Println("SortByPrice!")
	flights := getSampleFlights()

	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}
