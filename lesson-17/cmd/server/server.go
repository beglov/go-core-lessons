package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go-core-lessons/lesson-17/pkg/api"
	"log"
	"net/http"
	"os"
)

type server struct {
	api    *api.API
	router *mux.Router
}

func main() {
	srv := new(server)
	srv.router = mux.NewRouter()
	srv.api = api.New(srv.router)
	srv.api.Endpoints()
	loggedRouter := handlers.LoggingHandler(os.Stdout, srv.router)
	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}
