package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /oyle", oylesineHandler)
	mux.HandleFunc("GET /movie/{id}", getMovieHandler)
	mux.HandleFunc("POST /movie", postMovieHandler)
	fmt.Printf("Server is listening on port %s\n", server.Addr)
	server.ListenAndServe()
}
