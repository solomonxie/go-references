// run: go run hello_10_arrays.go
//
// An array has a fixed length baked into its type ([3]int is a different
// type from [4]int). Its size can't grow; that limitation is what the next
// step (slices) exists to solve.
//
// Step 10: arrays
package main

import "fmt"

func main() {
	var nums [3]int // Step 10: fixed-length array, zero-valued
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30

	days := [...]string{"Mon", "Tue", "Wed"} // Step 10: length inferred from literal

	fmt.Println(nums, len(nums))
	fmt.Println(days, len(days))

	for i, d := range days { // Step 10: range gives index + value
		fmt.Println(i, d)
	}
}
