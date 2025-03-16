package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
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
