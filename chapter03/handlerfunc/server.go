// In other words, anything that has a method called ServeHTTP with this method signature is a handler:
// ServeHTTP(http.ResponseWriter, *http.Request)
//
// We talked about handlers, but what are handler functions? Handler functions are functions that behave like handlers.
// Handler functions have the same signature as the ServeHTTP method;
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}

