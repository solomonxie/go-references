// run: go run hello_14_pointers.go
//
// & takes the address of a variable, producing a pointer (*T); * on a
// pointer dereferences it to read or write the value it points to. Passing
// a pointer to a function lets that function mutate the caller's variable.
//
// Step 14: pointers
package main

import "fmt"

type Point struct {
	X, Y int
}

func movePoint(p *Point, dx, dy int) { // Step 14: pointer parameter
	p.X += dx // Step 14: Go auto-dereferences p.X for (*p).X
	p.Y += dy
}

func main() {
	pt := Point{X: 1, Y: 1}
	ptr := &pt // Step 14: address-of

	fmt.Println("before:", pt, "ptr points to:", *ptr)

	movePoint(ptr, 5, 5)
	fmt.Println("after:", pt) // Step 14: mutation is visible to the caller
}
