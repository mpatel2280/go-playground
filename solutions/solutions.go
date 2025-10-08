package main

import (
	"fmt"

	"challenges/type_assertions"
)

func main() {
	fmt.Println("Type Assertions!")

	var name interface{} = "Abc"
	var age interface{} = 25

	dynamicDev := type_assertions.GetDeveloper(name, age)
	fmt.Println(dynamicDev.Name)
	fmt.Println(dynamicDev.Age)
}
