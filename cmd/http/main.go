package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)

	port := 3001

	log.Printf("Listening on port %d...", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	log.Fatal(err)
}
