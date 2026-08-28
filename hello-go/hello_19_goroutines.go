// run: go run hello_19_goroutines.go
//
// "go f()" starts f running concurrently as a goroutine instead of
// blocking the caller. main doesn't wait for goroutines by default, so a
// sync.WaitGroup is used here to block until all of them finish.
//
// Step 19: goroutines
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Step 19: signal completion when this goroutine returns
	fmt.Println("worker", id, "running")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)         // Step 19: register one goroutine to wait for
		go worker(i, &wg) // Step 19: launch concurrently
	}

	wg.Wait() // Step 19: block until all three call Done
	fmt.Println("all workers finished")
}
