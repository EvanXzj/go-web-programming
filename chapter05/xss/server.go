package main

import (
	"net/http"
	"html/template"
)

func process(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0") // 关闭浏览器xss攻击检测
	t, _ := template.ParseFiles("tmpl.html")
	// t.Execute(w, r.FormValue("comment"))
	t.Execute(w, template.HTML(r.FormValue("comment")))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/", form)
	server.ListenAndServe()
}
