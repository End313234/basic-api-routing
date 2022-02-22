package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

var router *chi.Mux = chi.NewRouter() // A new router that will be mounted into our main router

var methods map[string]func(pattern string, handlerFn http.HandlerFunc) = map[string]func(pattern string, handlerFn http.HandlerFunc){
	"GET":    router.Get,
	"POST":   router.Post,
	"PATCH":  router.Patch,
	"DELETE": router.Delete,
	"PUT":    router.Put,
}

type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

// Creates the route
func (route *Route) Create() {
	methods[route.Method](route.Pattern, route.Handler)
}

// Add a method
func (route *Route) AddMethod(method string) *Route {
	route.Method = method
	return route
}

// Add a pattern
func (route *Route) AddPattern(pattern string) *Route {
	route.Pattern = pattern
	return route
}

// Add a handler
func (route *Route) AddHandler(handler http.HandlerFunc) *Route {
	route.Handler = handler
	return route
}

// Just a function to improve readability when I get this router
func GetRouter() *chi.Mux {
	return router
}
