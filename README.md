# gitfinder
A web application that uses a golang backend to filter github repos and favorite them or delete them. Uses
the golang github api to do this.

To build the server:
```
go get -u github.com/gorilla/mux
go get -u github.com/google/go-github/github
go get -u golang.org/x/oauth2
go build server.go
```

<img src="https://github.com/snjt/gitfinder/blob/master/site.png" width="1195" height="634">
