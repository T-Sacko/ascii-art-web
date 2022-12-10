package main

import (
	"fmt"
	"html/template"
	//"log"xxd
	"net/http"
	//"os"
	//"strings"
	"gos/asciiart"
)

var tmpl *template.Template

func main() {
	//fmt.Println("Go")
	//tmpl, _ = template.ParseFiles("index.html")
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
	} else {
		tmpl, _ = template.ParseFiles("index.html")
	}
	fmt.Println(r.URL.Path)
	text := r.FormValue("message")
	font := r.FormValue("fonts")
	data := goss.AsciiReturn(font, text)
	tmpl.Execute(w, data)
}
// func AsciiReturn(f string, s string) string {
// 	var banner []byte
// 	var err error
// 	var answer string
// 	if s == "" {
// 		return ""
// 	}
// 	banner, err = os.ReadFile(f) // read file
// 	if err != nil {
// 		log.Fatal("ting messup")
// 	}
// 	//(asciinum-32)*9+1
// 	split := strings.Split(strings.ReplaceAll(string(banner), "\r", ""), "\n")
// 	st := strings.ReplaceAll(s, "\r", "") // seperates file by new lines
// 	myting := strings.Split(st, "\n")     // seperates the argument by '\n'
// 	for word := 0; word < len(myting); word++ {
// 		fmt.Println([]byte(myting[word]))
// 		//if word == 0 && len(myting) >= 3 {
// 		//	if len(myting[0]) == 0 && len(myting[1]) == 0 && len(myting[2]) == 0 {
// 		//	word += 1
// 		//}
// 		//} // current word
// 		for k := 0; k < 8; k++ { // every row
// 			for i := 0; i < len(myting[word]); i++ { // every letter in the word
// 				answer += split[(int(myting[word][i])-32)*9+1+k]
// 				// prints row of a letter
// 				// use += to add to "answer" variable
// 			}
// 			answer += "\n"

// 			if word == len(myting)-1 && len(myting[word]) == 0 {
// 				return answer
// 			}
// 		}
// 	}
// 	return answer
// }
