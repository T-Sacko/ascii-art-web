package main

import (
	"Gos"
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func main() {
	// fmt.Println("Go")
	// tmpl, _ = template.ParseFiles("index.html")
	http.HandleFunc("/", Handler)
	http.HandleFunc("/ascii", Ascii)
	http.ListenAndServe(":5505", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		tmpl, _ = template.ParseFiles("404.html")
	} else {
		tmpl, _ = template.ParseFiles("index.html")
	}

	fmt.Println(r.URL.Path)

	tmpl.Execute(w, nil)
}

func Ascii(w http.ResponseWriter, r *http.Request) {
	// http.NotFound(w, r)
	if r.URL.Path != "/ascii" {

		w.WriteHeader(http.StatusNotFound)
		tmpl, _ = template.ParseFiles("404.html")
	}
	fmt.Println(r.URL.Path)
	text := r.FormValue("message")
	font := r.FormValue("fonts")
	data := Gos.Asciii(font, text)
	if data == "no" {
		w.WriteHeader(http.StatusBadRequest)
		tmpl, _ = template.ParseFiles("400.html")
	} else if data == "nop" {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl, _ = template.ParseFiles("500.html")
	} else {
		tmpl, _ = template.ParseFiles("index.html")
	}

	// if data == "nop" {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	tmpl, _ = template.ParseFiles("500.html")
	// }
	tmpl.Execute(w, data)
}
