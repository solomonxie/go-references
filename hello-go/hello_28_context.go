// run: go run hello_28_context.go
//
// A context.Context carries a cancellation signal (and optional deadline)
// through a call chain. context.WithTimeout returns a context that cancels
// itself after the given duration; ctx.Done() is the channel that closes
// when that happens.
//
// Step 28: context
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	select {
	case <-time.After(500 * time.Millisecond): // Step 28: simulated slow work
		fmt.Println("worker finished")
	case <-ctx.Done(): // Step 28: context deadline wins the race
		fmt.Println("worker cancelled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel() // Step 28: always release the context's resources

	worker(ctx)
}
