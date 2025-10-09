package main

import (
	"challenges"
	"fmt"
	"slices"
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
// Deprecated: Use SortByPrice2 Or SortByPrice instead.
func SortByPrice1(flights []challenges.Flight) []challenges.Flight {
	// Simple Bubble sort or use sort.Slice for brevity
	// We'll use the standard library for cleaner code

	// imported := make([]challenges.Flight, len(flights))
	// copy(imported, flights)

	// Sort using sort.Slice
	// (you can also import "sort" at the top of the file)
	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Price < flights[j].Price
	})

	return flights
}

// SortByPrice sorts flights from highest to lowest
// Deprecated: Use SortByPrice instead.
func SortByPrice2(flights []challenges.Flight) []challenges.Flight {
	sorted := slices.Clone(flights)
	slices.SortFunc(sorted, func(a, b challenges.Flight) int {
		return a.Price - b.Price // negative = a < b
	})
	return sorted
}

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []challenges.Flight) []challenges.Flight {
	sorted := make([]challenges.Flight, len(flights))
	copy(sorted, flights)
	sort.Sort(challenges.ByPrice(sorted))
	return sorted
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
		{Origin: "Mumbai", Destination: "Dubai", Price: 250},
	}

	sortedList := SortByPrice(flights)
	printFlights(sortedList)

}
