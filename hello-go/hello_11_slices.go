// run: go run hello_11_slices.go
//
// A slice is a resizable view over an underlying array: append grows it,
// and slicing (s[low:high]) takes a sub-view without copying. Unlike an
// array, its type carries no length ([]int, not [3]int).
//
// Step 11: slices
package main

import "fmt"

func main() {
	nums := []int{10, 20, 30} // Step 11: slice literal, no length in the type
	nums = append(nums, 40)   // Step 11: append grows the slice

	fmt.Println(nums, "len:", len(nums), "cap:", cap(nums))

	middle := nums[1:3] // Step 11: sub-slice view, shares the backing array
	fmt.Println("middle:", middle)

	made := make([]int, 2, 5) // Step 11: make(type, len, cap)
	fmt.Println(made, "len:", len(made), "cap:", cap(made))
}
