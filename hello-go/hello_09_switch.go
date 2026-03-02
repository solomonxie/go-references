// run: go run hello_09_switch.go
//
// switch compares a value against several cases; unlike C, each case
// breaks automatically (no fallthrough unless the "fallthrough" keyword is
// used). A switch with no expression acts as a cleaner if/else-if chain.
//
// Step 9: switch
package main

import "fmt"

func dayName(d int) string {
	switch d { // Step 9: switch on a value
	case 0:
		return "Sunday"
	case 1:
		return "Monday"
	default:
		return "Other"
	}
}

func classify(n int) string {
	switch { // Step 9: no expression = if/else-if chain
	case n < 0:
		return "negative"
	case n == 0:
		return "zero"
	default:
		return "positive"
	}
}

func main() {
	fmt.Println(dayName(1), classify(-3))
}
