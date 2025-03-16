package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /oyle", oylesineHandler)
	mux.HandleFunc("GET /movie/{id}", getMovieHandler)
	mux.HandleFunc("POST /movie", postMovieHandler)
	fmt.Printf("Server is listening on port %s\n", server.Addr)
	server.ListenAndServe()
}
