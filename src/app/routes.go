package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mpludowski/bWljaGFsLXBsdWRvd3NraQo/controller"
)

func setupRoutes() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.Methods("POST").Path("/fetcher").HandlerFunc(sizeLimitMiddleware(controller.PostFetcher))
	api.Methods("DELETE").Path("/fetcher/{id}").HandlerFunc(sizeLimitMiddleware(controller.DeleteFetcher))
	api.Methods("GET").Path("/fetcher").HandlerFunc(sizeLimitMiddleware(controller.ListFetchers))
	api.Methods("GET").Path("/fetcher/{id}/history").HandlerFunc(sizeLimitMiddleware(controller.History))

	return r
}

func sizeLimitMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1*1024*1024) // 1 MB
		f(w, r)
	}
}
