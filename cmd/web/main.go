package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./assets/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}/", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Started server at :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
