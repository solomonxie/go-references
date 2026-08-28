// run: go run hello_27_mutex.go
//
// Multiple goroutines writing the same variable at once is a data race.
// sync.Mutex serializes access: Lock blocks until no other goroutine holds
// the lock, Unlock releases it. defer pairs the two reliably.
//
// Step 27: sync.Mutex
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex // Step 27: guards value
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock() // Step 27: always released, even if Inc panicked
	c.value++
}

func main() {
	c := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ { // Step 27: 100 goroutines racing to increment
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(c.value) // Step 27: always 100, thanks to the mutex
}
