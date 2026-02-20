// run: go run hello_06_named_returns.go
//
// Return values can be named in the function signature. A named return
// acts as a pre-declared local variable; a bare "return" (a "naked
// return") sends back whatever those variables currently hold.
//
// Step 6: named return values
package main

import "fmt"

func divmod(a, b int) (q, r int) { // Step 6: q, r declared by the signature
	q = a / b
	r = a % b
	return // Step 6: naked return sends back q, r
}

func main() {
	q, r := divmod(17, 5)
	fmt.Println("quotient:", q, "remainder:", r)
}
