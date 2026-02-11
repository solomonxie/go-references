// run: go run hello_03_constants.go
//
// const declares a value fixed at compile time. iota is a counter that
// resets to 0 in each const block and increments per line, used to build
// enum-like sequences without writing each number by hand.
//
// Step 3: constants
package main

import "fmt"

const Pi = 3.14159

const (
	Sunday = iota // Step 3: iota starts at 0
	Monday
	Tuesday
	Wednesday
)

func main() {
	fmt.Println("Pi:", Pi)
	fmt.Println("Wednesday index:", Wednesday)
}
