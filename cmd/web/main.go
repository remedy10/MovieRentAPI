package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

var logger *slog.Logger

func main() {
	addr := flag.String("addr", ":8080", "HTTP Network Address")
	flag.Parse()
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    *addr,
		Handler: mux,
	}
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /oyle", oylesineHandler)
	mux.HandleFunc("GET /movie/{id}", getMovieHandler)
	mux.HandleFunc("POST /movie", postMovieHandler)
	logger.Info("Server is listening on port", slog.String("addr", server.Addr))
	server.ListenAndServe()
}
