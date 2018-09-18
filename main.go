package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", handleIndex)
	fmt.Println("[Portofolio App] Running ...")
	http.ListenAndServe(":8080", nil)
}

func handleIndex(wr http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(wr, "index.gohtml", nil)
}
