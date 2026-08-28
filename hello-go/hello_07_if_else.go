// run: go run hello_07_if_else.go
//
// if/else needs no parentheses around the condition and always uses
// braces. An if can carry an init statement before the condition, scoped
// to the if/else chain only.
//
// Step 7: if / else
package main

import "fmt"

func classify(n int) string {
	if n < 0 { // Step 7: no parens around condition
		return "negative"
	} else if n == 0 {
		return "zero"
	} else {
		return "positive"
	}
}

func main() {
	fmt.Println(classify(-5), classify(0), classify(5))

	if half := 10 / 2; half > 3 { // Step 7: init statement scoped to this if/else
		fmt.Println("half is", half)
	}
}
