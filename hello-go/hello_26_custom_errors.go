// run: go run hello_26_custom_errors.go
//
// Any type with an Error() string method satisfies the error interface, so
// a custom struct can carry structured data alongside its message.
// errors.As recovers that concrete type back out of an error value.
//
// Step 26: custom error types
package main

import (
	"errors"
	"fmt"
)

type ValidationError struct { // Step 26: struct carrying structured error data
	Field string
	Msg   string
}

func (e *ValidationError) Error() string { // Step 26: satisfies the error interface
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

func validate(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Msg: "must not be negative"}
	}
	return nil
}

func main() {
	err := validate(-1)
	fmt.Println(err)

	var ve *ValidationError
	if errors.As(err, &ve) { // Step 26: recover the concrete *ValidationError
		fmt.Println("field:", ve.Field)
	}
}
