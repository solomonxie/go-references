// run: go run hello_23_variadic.go
//
// A parameter written as ...T accepts any number of trailing arguments,
// collected inside the function as a []T. A slice can be "spread" into
// such a parameter with the same ... suffix at the call site.
//
// Step 23: variadic functions
package main

import "fmt"

func sum(nums ...int) int { // Step 23: nums is a []int inside the function
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))

	values := []int{4, 5, 6}
	fmt.Println(sum(values...)) // Step 23: spread a slice into the variadic call
}
