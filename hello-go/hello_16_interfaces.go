// run: go run hello_16_interfaces.go
//
// An interface lists method signatures; any type that implements them
// satisfies the interface implicitly (no "implements" keyword). A variable
// of interface type can hold any concrete type that qualifies.
//
// Step 16: interfaces
package main

import "fmt"

type Shape interface { // Step 16: interface = method set
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 { // Step 16: Circle satisfies Shape implicitly
	return 3.14159 * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 { // Step 16: Rectangle also satisfies Shape
	return r.Width * r.Height
}

func describe(s Shape) { // Step 16: accepts any Shape
	fmt.Printf("area: %.2f\n", s.Area())
}

func main() {
	describe(Circle{Radius: 2})
	describe(Rectangle{Width: 3, Height: 4})
}
