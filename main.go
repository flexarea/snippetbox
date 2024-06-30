package main

import (
	"log"
	"net/http"
)

// home handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from snippetbox"))
}

func main() {
	newmux := http.NewServeMux()
	newmux.HandleFunc("/", home)

	log.Println("Starting server on: 4000")
	err := http.ListenAndServe(":4000", newmux)
	log.Fatal(err)
}
