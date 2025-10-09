package main

import (
	"challenges"
	"fmt"
	"slices"
	"sort"
)

// SortByPrice sorts flights from highest to lowest using the challenges.ByPrice type
func SortByPrice(flights []challenges.Flight) []challenges.Flight {
	sorted := make([]challenges.Flight, len(flights))
	copy(sorted, flights)
	sort.Sort(challenges.ByPrice(sorted))
	return sorted
}

// SortByPrice1 sorts flights from highest to lowest using sort.Slice
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

// SortByPrice2 sorts flights from highest to lowest using slices.SortFunc
// Deprecated: Use SortByPrice instead.
func SortByPrice2(flights []challenges.Flight) []challenges.Flight {
	sorted := slices.Clone(flights)
	slices.SortFunc(sorted, func(a, b challenges.Flight) int {
		return a.Price - b.Price // negative = a < b
	})
	return sorted
}

// getSampleFlights returns a sample slice of flights for demonstration
func getSampleFlights() []challenges.Flight {
	return []challenges.Flight{
		{Origin: "New York", Destination: "London", Price: 800},
		{Origin: "Paris", Destination: "Berlin", Price: 120},
		{Origin: "Tokyo", Destination: "Sydney", Price: 950},
		{Origin: "Delhi", Destination: "Dubai", Price: 300},
		{Origin: "Mumbai", Destination: "Dubai", Price: 250},
	}
}

// printFlights prints flight information in a formatted way
func printFlights(flights []challenges.Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d\n", flight.Origin, flight.Destination, flight.Price)
	}
}
