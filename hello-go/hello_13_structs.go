// run: go run hello_13_structs.go
//
// A struct groups named fields into one type. Struct literals can set
// fields by name (order-independent) or positionally (must match field
// order exactly).
//
// Step 13: structs
package main

import "fmt"

type Point struct { // Step 13: struct type with two fields
	X, Y int
}

func main() {
	p1 := Point{X: 1, Y: 2} // Step 13: named-field literal
	p2 := Point{3, 4}       // Step 13: positional literal

	fmt.Println(p1, p2)
	fmt.Println("p1.X:", p1.X) // Step 13: dot access

	p1.Y = 20 // Step 13: fields are mutable
	fmt.Println(p1)
}
