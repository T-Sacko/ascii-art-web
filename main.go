package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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
	data := font
	data = AsciiReturn(font, text)
	tmpl.Execute(w, data)
}

func AsciiReturn(f string, s string) string {
	var banner []byte
	var err error
	var answer string

	banner, err = os.ReadFile(f) // read file

	if err != nil {
		log.Fatal("ting messup")
	}
	//(asciinum-32)*9+1
	split := strings.Split(string(banner), "\n") // seperates file by new lines
	myting := strings.Split(s, "\\n")            // seperates the argument by '\n'
	for word := 0; word < len(myting); word++ {

		if word == 0 && len(myting) >= 3 {
			if len(myting[0]) == 0 && len(myting[1]) == 0 && len(myting[2]) == 0 {
				word += 1
			}
		} // current word
		for k := 0; k < 8; k++ { // every row
			if len(myting[word]) == 0 && len(myting) >= 2 { // checks for '\n'
				k = 7
			}
			for i := 0; i < len(myting[word]); i++ { // every letter in the word
				answer += split[(int(myting[word][i])-32)*9+1+k] // prints row of a letter
				// use += to add to "answer" variable
			}
			if len(myting[word]) != 0 {
				answer += "\n"
				// use += to add "\n" to "answer" variable
			}
			if len(myting[word]) == 0 && len(myting) >= 2 {
				fmt.Println()
				if len(myting) == 2 && word != len(myting)-1 {
					if len(myting[word+1]) == 0 {
						word++
					}
				}
			}
		}
	}
	return answer
}
