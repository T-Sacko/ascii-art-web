package main

import (
	"net/http"
	"text/template"
)

var tmpl *template.Template

type Data struct {
	n int
	s string
}

func main() {
	tmpl, _ = template.ParseFiles("index.html")
	http.HandleFunc("/", Handler)
	http.HandleFunc("/ascii", Ascii)
	http.ListenAndServe("localhost:8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func Ascii(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("message")
	font := r.FormValue("fonts")
	data := ""
	data = AsciiReturn(font, text)
	tmpl.Execute(w, data)
}
