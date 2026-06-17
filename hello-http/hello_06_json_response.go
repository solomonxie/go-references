// run: go run hello_06_json_response.go
// then: curl http://localhost:8086/status
//
// json.NewEncoder(w).Encode streams a Go value out as JSON directly onto
// the response writer. The Content-Type header is set by hand since Go
// doesn't infer it automatically.
//
// Step 6: encoding a JSON response
package main

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	OK bool `json:"ok"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") // Step 6: declare the body type
		json.NewEncoder(w).Encode(Status{OK: true})        // Step 6: struct -> JSON on the wire
	})

	http.ListenAndServe(":8086", mux)
}
