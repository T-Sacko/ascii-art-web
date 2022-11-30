package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe("localhost:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello guys")
}
