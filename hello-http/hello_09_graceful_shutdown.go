// run: go run hello_09_graceful_shutdown.go
// then: curl http://localhost:8089/, and press Ctrl+C in the server terminal to shut it down
//
// An *http.Server (instead of the package-level ListenAndServe) can be
// shut down deliberately: Shutdown stops accepting new connections and
// waits for in-flight ones to finish, bounded by a context deadline.
//
// Step 9: graceful shutdown
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	srv := &http.Server{Addr: ":8089", Handler: mux} // Step 9: explicit server, not the package default

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt) // Step 9: catch Ctrl+C
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx) // Step 9: stop accepting, drain in-flight requests
	log.Println("server stopped")
}
