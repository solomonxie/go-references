// run: go run hello_17_errors.go
//
// Go has no exceptions for ordinary failures: functions return an error as
// their last value, and callers check it with "if err != nil". errors.New
// and fmt.Errorf build error values; %w wraps an underlying error so it
// can be unwrapped later.
//
// Step 17: errors
package main

import (
	"errors"
	"fmt"
)

var ErrDivByZero = errors.New("division by zero") // Step 17: sentinel error value

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

func compute(a, b float64) (float64, error) {
	result, err := divide(a, b)
	if err != nil {
		return 0, fmt.Errorf("compute failed: %w", err) // Step 17: wrap with context
	}
	return result, nil
}

func main() {
	if result, err := compute(10, 2); err == nil {
		fmt.Println("result:", result)
	}

	_, err := compute(10, 0)
	if err != nil {
		fmt.Println("error:", err)
		fmt.Println("is div-by-zero:", errors.Is(err, ErrDivByZero)) // Step 17: unwrap to compare
	}
}
