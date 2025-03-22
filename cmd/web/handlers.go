package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/pages/home.tmpl",
		"./ui/html/pages/base.tmpl",
		"./ui/html/partials/nav.tmpl",
	}
	tmp, err := template.ParseFiles(files...)
	if err != nil {
		logger.Error("an error executing template file ", slog.String("err", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmp.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("an error executing template file ", slog.String("err", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	logger.Info("request successful", slog.String("method", r.Method), slog.String("path", r.URL.Path))
}

func oylesineHandler(w http.ResponseWriter, r *http.Request) {
	w.Header()["saNaNe-La"] = []string{"ne ddiyon La BeBe"}
	w.Write([]byte("oylesine handler"))
}

func getMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 0 {
		logger.Error("invalid id", slog.String("err", err.Error()))
		http.NotFound(w, r)
		return
	}
	logger.Info("get movie handler", slog.String("method", r.Method), slog.String("path", r.URL.Path))
	fmt.Fprintf(w, "get movie handler %s", r.PathValue("id"))
}
func postMovieHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("post movie handler", slog.String("method", r.Method), slog.String("path", r.URL.Path))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("post movie handler"))
}
