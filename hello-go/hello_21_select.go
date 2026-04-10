// run: go run hello_21_select.go
//
// select waits on several channel operations at once and runs whichever
// case is ready first, like a switch for channels. A default case makes it
// non-blocking; time.After produces a channel that fires after a delay,
// commonly used as a timeout case.
//
// Step 21: select
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "from ch1"
	}()
	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	for i := 0; i < 2; i++ {
		select { // Step 21: runs whichever channel is ready first
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

	select { // Step 21: default makes this a non-blocking check
	case msg := <-ch1:
		fmt.Println(msg)
	default:
		fmt.Println("no message ready")
	}

	select { // Step 21: time.After as a timeout case
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(30 * time.Millisecond):
		fmt.Println("timed out waiting for ch1")
	}
}
