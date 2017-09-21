package main

import (
	"github.com/Grazfather/TheGoProgrammingLanguage/ch1/lissajous"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL.
func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	lissajous.Lissajous(w, r.Form)
}
