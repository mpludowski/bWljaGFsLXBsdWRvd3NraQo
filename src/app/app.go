package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	Name   string
}

func NewApp(name string) App {
	a := App{}
	a.Name = name
	a.Router = setupRoutes()

	return a
}

func (a *App) Run(addr string) {
	log.Print("Starting server...")

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         addr,
		Handler:      a.Router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
