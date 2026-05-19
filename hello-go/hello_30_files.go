// run: go run hello_30_files.go
//
// os.CreateTemp makes a uniquely named file so this step never collides
// with another run or another step. Writing and reading a real file go
// through the os package rather than an in-memory buffer.
//
// Step 30: file I/O
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.CreateTemp("", "hello_30_*.txt") // Step 30: unique temp file
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name()) // Step 30: cleanup; runs after Close (LIFO)
	defer f.Close()

	if _, err := f.WriteString("hello from go\n"); err != nil {
		panic(err)
	}

	data, err := os.ReadFile(f.Name()) // Step 30: independent read from disk
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
}
