// run: go run hello_18_defer_panic_recover.go
//
// defer schedules a call to run when the enclosing function returns (LIFO
// order), used for cleanup regardless of how the function exits. panic
// aborts normal execution; recover, called inside a deferred function,
// stops the panic and lets the program continue.
//
// Step 18: defer, panic, recover
package main

import "fmt"

func safeDivide(a, b int) (result int, err error) {
	defer func() { // Step 18: recover runs only inside a deferred call
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()

	result = a / b // Step 18: dividing by zero panics here
	return
}

func main() {
	defer fmt.Println("main: cleanup runs last") // Step 18: deferred, runs on return
	fmt.Println("main: start")

	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", result)
	}
}
