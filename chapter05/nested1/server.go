package main

import (
	"net/http"
	"html/template"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html")
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
