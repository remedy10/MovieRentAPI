package main

import (
	"fmt"
	"html/template"
	"log"
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
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmp.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print("an error executing template file ", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
}

func oylesineHandler(w http.ResponseWriter, r *http.Request) {
	w.Header()["saNaNe-La"] = []string{"ne ddiyon La BeBe"}
	w.Write([]byte("oylesine handler"))
}

func getMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "get movie handler %s", r.PathValue("id"))
}
func postMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("post movie handler"))
}
