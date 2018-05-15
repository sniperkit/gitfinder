package main

import (
	"gitfilter/resources"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const URL string = "localhost"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/repos/{filter}", resources.GitCliHandler(resources.FilterHandler)).
		Methods("GET")

	r.HandleFunc("/favorites", resources.NewFavoriteHandler).
		Methods("POST")

	r.HandleFunc("/favorites", resources.GitCliHandler(resources.GetFavoritesHandler)).
		Methods("GET")

	r.Use(resources.CORS)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
