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

	fmt.Println("Filter Unique Challenge")
	users := getSampleUsers()

	uniqueNames := FilterUnique(users)
	fmt.Println(uniqueNames)

	freqCheck()

	fmt.Println("map and structs!")
	mapStructs()
}

func freqCheck() {
	words := []string{"go", "java", "go", "python", "go"}
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}
	fmt.Println(freq) // map[go:3 java:1 python:1]

}

func mapStructs() {
	var m map[string]challenges.User = make(map[string]challenges.User)
	m["one"] = challenges.User{Name: "Alice", Age: 25}
	m["two"] = challenges.User{Name: "Bob", Age: 30}
	fmt.Println(m) // Output: map[one:{Alice 25} two:{Bob 30}]
}
