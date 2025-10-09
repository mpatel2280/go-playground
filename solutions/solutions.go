package main

import (
	"challenges"
	"fmt"
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

	getEngg(dynamicDev.Name, dob)

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
