package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	ae := r.Header["Accept-Encoding"] // map of strings
	fmt.Printf("%T - %v\n", ae, ae)
	ae2 := r.Header.Get("Accept-Encoding") // string类型
	fmt.Printf("%T - %v\n", ae2, ae2)
	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
