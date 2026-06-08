// run: go run hello_03_path_params.go
// then: curl http://localhost:8083/users/42
//
// A {name} segment in a registered pattern captures part of the path;
// r.PathValue reads it back out by name inside the handler.
//
// Step 3: path parameters
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) { // Step 3: {id} segment
		id := r.PathValue("id") // Step 3: read the captured segment
		fmt.Fprintf(w, "user id: %s\n", id)
	})

	http.ListenAndServe(":8083", mux)
}
