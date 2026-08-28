// run: go run hello_24_type_assertions.go
//
// any (an alias for interface{}) can hold a value of any type. A type
// assertion (v.(T)) recovers the concrete type back out; the comma-ok form
// avoids a panic when the assertion fails. A type switch handles several
// possible concrete types in one place.
//
// Step 24: type assertions and type switches
package main

import "fmt"

func describe(v any) string {
	switch x := v.(type) { // Step 24: type switch on the dynamic type
	case int:
		return fmt.Sprintf("int: %d", x)
	case string:
		return fmt.Sprintf("string: %q", x)
	default:
		return fmt.Sprintf("other: %v", x)
	}
}

func main() {
	fmt.Println(describe(42))
	fmt.Println(describe("hi"))

	var v any = "direct"
	s, ok := v.(string) // Step 24: comma-ok assertion, no panic on mismatch
	fmt.Println(s, ok)

	n, ok := v.(int) // Step 24: wrong type: zero value + false, not a panic
	fmt.Println(n, ok)
}
