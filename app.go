package main

import "net/http"
import "html/template"

func main() {
	http.HandleFunc("/", indexAction)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

var templates = template.Must(template.ParseGlob("views/*"))

func indexAction(rw http.ResponseWriter, req *http.Request) {
	// rw.Write([]byte("Hello World"))

	err := templates.ExecuteTemplate(rw, "index.html", nil)

	if err != nil {
		rw.WriteHeader(404)
		rw.Write([]byte("Page Not Found"))
	}
}
