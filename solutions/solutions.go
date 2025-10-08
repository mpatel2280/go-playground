package main

import (
	"fmt"

	"challenges/type_assertions"
)

func main() {
	fmt.Println("Type Assertions!")

	var name interface{} = "Abcd"
	var age interface{} = 24

	dynamicDev := type_assertions.GetDeveloper(name, age)
	fmt.Println(dynamicDev.Name)
	fmt.Println(dynamicDev.Age)
}
