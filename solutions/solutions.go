package main

import (
	"challenges"
	"fmt"
	"sort"
	"time"
)

type Engineer struct {
	Name string
	DOB  time.Time
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

func (e Engineer) Age() int {
	today := time.Now()
	years := today.Year() - e.DOB.Year()

	// Adjust if birthday hasn't occurred yet this year
	if today.YearDay() < e.DOB.YearDay() {
		years--
	}

	return years
}

func getEngg(name string, dob time.Time) Engineer {
	engg := Engineer{Name: name, DOB: dob}

	var programmers []challenges.Employee
	programmers = append(programmers, engg)

	for _, p := range programmers {
		fmt.Printf("%s and is %d years old.\n", p.Language(), p.Age())
	}
	return engg
}

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []challenges.Flight) []challenges.Flight {
	// Simple Bubble sort or use sort.Slice for brevity
	// We'll use the standard library for cleaner code
	imported := make([]challenges.Flight, len(flights))
	copy(imported, flights)

	// Sort using sort.Slice
	// (you can also import "sort" at the top of the file)
	sort.Slice(imported, func(i, j int) bool {
		return imported[i].Price < imported[j].Price
	})

	return imported
}

func printFlights(flights []challenges.Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d\n", flight.Origin, flight.Destination, flight.Price)
	}
}

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
	flights := []challenges.Flight{
		{Origin: "New York", Destination: "London", Price: 800},
		{Origin: "Paris", Destination: "Berlin", Price: 120},
		{Origin: "Tokyo", Destination: "Sydney", Price: 950},
		{Origin: "Delhi", Destination: "Dubai", Price: 300},
	}

	sortedList := SortByPrice(flights)
	printFlights(sortedList)

}
