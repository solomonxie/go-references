// run: go run hello_32_cli_flags.go -name=Go
//
// The flag package parses command-line arguments into typed variables.
// flag.String registers a -name flag with a default; flag.Parse must run
// before any flag value is read.
//
// Step 32: CLI flags
package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "world", "name to greet") // Step 32: -name flag, default "world"
	flag.Parse()                                          // Step 32: must run before *name is read

	fmt.Printf("Hello, %s!\n", *name)
}
