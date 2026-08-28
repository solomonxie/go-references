// run: go run hello_20_channels.go
//
// A channel is a typed pipe goroutines use to send and receive values
// safely, replacing manual locks for this kind of handoff. An unbuffered
// channel (make(chan T)) blocks the sender until a receiver is ready, so
// send and receive rendezvous.
//
// Step 20: channels
package main

import "fmt"

func worker(id int, results chan<- string) { // Step 20: send-only channel type
	results <- fmt.Sprintf("worker %d done", id) // Step 20: blocks until received
}

func main() {
	results := make(chan string) // Step 20: unbuffered channel

	for i := 1; i <= 3; i++ {
		go worker(i, results)
	}

	for i := 0; i < 3; i++ {
		msg := <-results // Step 20: receive blocks until a value arrives
		fmt.Println(msg)
	}
}
