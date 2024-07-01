package main

import (
	"log"
	"net/http"
)

// home handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from snippetbox"))
}
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))
}
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/createSnippet", showSnippet)

	log.Println("Starting server on: 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
