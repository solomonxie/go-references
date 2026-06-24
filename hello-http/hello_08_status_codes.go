// run: go run hello_08_status_codes.go
// then: curl -i http://localhost:8088/users/1   and   curl -i http://localhost:8088/users/9
//
// http.Error writes a status code plus a plain-text body in one call;
// w.WriteHeader sets the code explicitly when the body is written by hand
// instead. The header must be set before any call to w.Write.
//
// Step 8: status codes
package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id != "1" {
			http.Error(w, "not found", http.StatusNotFound) // Step 8: 404 + message
			return
		}
		w.WriteHeader(http.StatusOK) // Step 8: explicit 200 before writing the body
		w.Write([]byte("found user 1"))
	})

	http.ListenAndServe(":8088", mux)
}
