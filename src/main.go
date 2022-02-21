package main

import (
	"net/http"

	_ "github.com/End313234/basic-api-routing/src/api"
	"github.com/End313234/basic-api-routing/src/handlers"
	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	routes := handlers.GetRouter()

	router.Mount("/", routes)

	http.ListenAndServe(":3000", router)
}
