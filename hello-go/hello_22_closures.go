// run: go run hello_22_closures.go
//
// A function literal defined inside another function captures the
// enclosing variables by reference, not by copy. Each call to makeCounter
// gets its own independent count.
//
// Step 22: closures
package main

import "fmt"

func makeCounter() func() int {
	count := 0
	return func() int { // Step 22: captures count from the enclosing scope
		count++
		return count
	}
}

func main() {
	next := makeCounter()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	other := makeCounter() // Step 22: independent count, unaffected by next
	fmt.Println(other())
}
