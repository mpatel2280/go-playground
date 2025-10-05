package main

import (
	"fmt"
	greetings "module-example/greetings"
)

func main() {
	message := greetings.Hello("World")
	fmt.Println(message)
}
