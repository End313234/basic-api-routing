package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

var router *chi.Mux = chi.NewRouter() // A new router that will be mounted into our main router

// Creating a route
func CreateRoute(handler http.Handler) {
	router.Mount("/", handler)
}

// Just a function to improve readability when I get this router
func GetRouter() *chi.Mux {
	return router
}
