package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

// Ref
// https://go.dev/doc/tutorial/create-module
// https://go.dev/doc/tutorial/call-module-code
// https://go.dev/doc/tutorial/handle-errors
