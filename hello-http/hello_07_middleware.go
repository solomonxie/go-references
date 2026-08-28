// run: go run hello_07_middleware.go
// then: curl http://localhost:8087/  (watch the log line in the server terminal)
//
// A middleware wraps an http.Handler in another http.Handler that runs
// code before/after calling the original — here, logging how long each
// request took. Handlers compose by wrapping, with no framework involved.
//
// Step 7: middleware
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func logging(next http.Handler) http.Handler { // Step 7: wraps a handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r) // Step 7: delegate to the wrapped handler
		log.Printf("%s %s took %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "home")
	})

	http.ListenAndServe(":8087", logging(mux)) // Step 7: mux wrapped by logging
}
