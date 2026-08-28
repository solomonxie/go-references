// run: go run hello_04_query_params.go
// then: curl "http://localhost:8084/search?q=go"
//
// r.URL.Query() parses the "?key=value" part of the URL into a map-like
// Values type; Get returns "" if the key wasn't present.
//
// Step 4: query parameters
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q") // Step 4: read ?q= from the query string
		fmt.Fprintf(w, "searching for: %s\n", q)
	})

	http.ListenAndServe(":8084", mux)
}
