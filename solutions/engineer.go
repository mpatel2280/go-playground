package main

import (
	"challenges"
	"fmt"
	"time"
)

// Engineer represents a software engineer with name and date of birth
type Engineer struct {
	Name string
	DOB  time.Time
}

// Language returns a string describing the programming language the engineer uses
func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

// Age calculates and returns the current age of the engineer
func (e Engineer) Age() int {
	today := time.Now()
	years := today.Year() - e.DOB.Year()

	// Adjust if birthday hasn't occurred yet this year
	if today.YearDay() < e.DOB.YearDay() {
		years--
	}

	return years
}

// getEngg creates an Engineer instance and demonstrates the Employee interface
func getEngg(name string, dob time.Time) Engineer {
	engg := Engineer{Name: name, DOB: dob}

	var programmers []challenges.Employee
	programmers = append(programmers, engg)

	for _, p := range programmers {
		fmt.Printf("%s and is %d years old.\n", p.Language(), p.Age())
	}
	return engg
}
