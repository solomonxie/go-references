// run: go run hello_02_routing.go
// then: curl http://localhost:8082/  and  curl http://localhost:8082/about
//
// http.NewServeMux creates a router explicitly instead of relying on the
// package-level default mux, so each route is registered on it by hand.
//
// Step 2: routing
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // Step 2: explicit router

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "home")
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "about page")
	})

	http.ListenAndServe(":8082", mux)
}
