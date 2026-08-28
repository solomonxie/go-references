// run: go run hello_01_hello_server.go
// then, in another terminal: curl http://localhost:8081/
//
// http.HandleFunc registers a handler for a path on the default mux;
// http.ListenAndServe starts the server and blocks forever, calling that
// handler once per request.
//
// Step 1: minimal HTTP server
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // Step 1: handler for "/"
		fmt.Fprintln(w, "Hello, Go!")
	})

	http.ListenAndServe(":8081", nil) // Step 1: blocks, serving forever
}
