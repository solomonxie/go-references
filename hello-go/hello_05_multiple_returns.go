// run: go run hello_05_multiple_returns.go
//
// A Go function can return more than one value. This is the idiomatic way
// to return a result alongside an error or a secondary piece of data,
// instead of packing them into a struct or tuple.
//
// Step 5: multiple return values
package main

import "fmt"

func divmod(a, b int) (int, int) { // Step 5: two return types
	return a / b, a % b
}

func main() {
	q, r := divmod(17, 5) // Step 5: both results captured at once
	fmt.Println("quotient:", q, "remainder:", r)
}
