package resources

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"net/http"
)

const API_KEY string = "5bac038dd30f3a3c11001ce34042eba331d05746"

var gitClient *github.Client = nil
var ctx *context.Context = nil

func getClient() (*github.Client, *context.Context) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: API_KEY})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client, &ctx
}

func CORS(handle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		handle.ServeHTTP(w, r)
	})
}

func GitCliHandler(handler func(cli *github.Client,
	ctx *context.Context,
	w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	cli, ctx := getClient()
	return func(w http.ResponseWriter, r *http.Request) {
		handler(cli, ctx, w, r)
	}
}
