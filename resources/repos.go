package resources

import (
	"context"
	"encoding/json"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func FilterHandler(cli *github.Client, ctx *context.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var filter string = vars["filter"]
	log.Println(filter)
	opts := &github.SearchOptions{Sort: "stars", Order: "desc"}
	repos, _, _ := cli.Search.Repositories(*ctx, filter, opts)
	var returnVal map[string][]github.Repository = make(map[string][]github.Repository)
	returnVal["repositories"] = repos.Repositories
	jsonData, _ := json.Marshal(returnVal)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
