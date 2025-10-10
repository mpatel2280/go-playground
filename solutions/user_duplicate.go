package main

import (
	"challenges"
)

// FilterUnique returns a slice of unique user names
func FilterUnique(users []challenges.User) []string {
	unique := make(map[string]bool)
	for _, user := range users {
		unique[user.Name] = true
	}
	var uniqueNames []string
	for name := range unique {
		uniqueNames = append(uniqueNames, name)
	}
	return uniqueNames
}

// getSampleUsers returns a sample slice of users for demonstration
func getSampleUsers() []challenges.User {
	return []challenges.User{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Alice", Age: 25},
		{Name: "Charlie", Age: 35},
	}
}
