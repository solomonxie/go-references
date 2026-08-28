// run: go run hello_08_for_loops.go
//
// for is Go's only loop keyword; it covers what other languages split into
// for/while/do. Three shapes: classic init;cond;post, condition-only
// (a "while"), and no condition at all (infinite, exited with break).
//
// Step 8: for loops
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ { // Step 8: classic three-part for
		fmt.Println("classic", i)
	}

	n := 0
	for n < 3 { // Step 8: condition-only, acts like "while"
		fmt.Println("while-style", n)
		n++
	}

	c := 0
	for { // Step 8: infinite loop, exited explicitly
		if c >= 3 {
			break
		}
		fmt.Println("infinite", c)
		c++
	}
}
