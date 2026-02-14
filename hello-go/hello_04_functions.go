// run: go run hello_04_functions.go
//
// func declares a function: name, parameters with their types, and a return
// type. Parameters of the same type can share one type annotation.
//
// Step 4: functions
package main

import "fmt"

func add(a, b int) int { // Step 4: shared type for a, b; single return type
	return a + b
}

func greet(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(greet("Go"))
}
