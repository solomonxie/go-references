// run: go run hello_15_methods.go
//
// A method is a function with a receiver: func (r Type) Name(...). A value
// receiver gets a copy and can't mutate the original; a pointer receiver
// (*Type) operates on the original and is how methods mutate state.
//
// Step 15: methods
package main

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) String() string { // Step 15: value receiver, read-only
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p *Point) Move(dx, dy int) { // Step 15: pointer receiver, mutates
	p.X += dx
	p.Y += dy
}

func main() {
	pt := Point{X: 1, Y: 1}
	fmt.Println(pt.String())

	pt.Move(4, 4) // Step 15: Go auto-takes &pt to call the pointer-receiver method
	fmt.Println(pt.String())
}
