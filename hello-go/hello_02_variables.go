// run: go run hello_02_variables.go
//
// var declares a typed variable, optionally with an initial value; Go infers
// the type when one is given. := is short-hand for "var + infer" and only
// works inside a function body. A var block groups several declarations.
//
// Step 2: variables
package main

import "fmt"

func main() {
	var name string = "Gopher" // explicit type
	var age = 5                // inferred: int
	count := 3                 // Step 2: := infers type and declares in one step

	var (
		x = 1
		y = 2
	)

	fmt.Println(name, age, count, x, y)
}
