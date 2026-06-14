// run: go run hello_05_json_request.go
// then: curl -X POST -d '{"name":"Go"}' http://localhost:8085/greet
//
// json.NewDecoder(r.Body).Decode streams the request body straight into a
// Go struct, without reading the whole body into memory first.
//
// Step 5: decoding a JSON request body
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Greeting struct {
	Name string `json:"name"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		var g Greeting
		if err := json.NewDecoder(r.Body).Decode(&g); err != nil { // Step 5: body -> struct
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello, %s!\n", g.Name)
	})

	http.ListenAndServe(":8085", mux)
}
