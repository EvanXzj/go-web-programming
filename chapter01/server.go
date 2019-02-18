package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s", req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running at :3000")
	http.ListenAndServe(":3000", nil)
}
