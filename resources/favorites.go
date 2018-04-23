package resources

import (
	"context"
	"encoding/json"
	"github.com/google/go-github/github"
	//"log"
	"net/http"
)

type getFavorites func([]string) []github.Repository
type addFavorites func([]string)

var favoriteRepos []github.Repository = []github.Repository{}

type Favorites struct {
	Repos []github.Repository
}

//func NewFavoriteHandler(fn addFavorites, w http.ResponseWriter, r *http.Request) {
func NewFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	var favorites Favorites
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&favorites)
	if err != nil {
		panic(err)
	}
	favoriteRepos = append(favoriteRepos, favorites.Repos...)
	defer r.Body.Close()
	var returnVal map[string]string = make(map[string]string)
	returnVal["message"] = "Favorites added."
	jsonData, _ := json.Marshal(returnVal)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

//func GetFavoritesHandler(fn getFavorites, w http.ResponseWriter, r *http.Request) {
func GetFavoritesHandler(cli *github.Client, ctx *context.Context, w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	var returnVal map[string][]github.Repository = make(map[string][]github.Repository)
	returnVal["repositories"] = favoriteRepos
	jsonData, _ := json.Marshal(returnVal)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
