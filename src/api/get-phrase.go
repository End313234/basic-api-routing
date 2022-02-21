package api

import (
	"io/ioutil"
	"net/http"

	"github.com/End313234/basic-api-routing/src/handlers"
	"github.com/go-chi/chi"
)

func init() {
	// Instanciating the router for this file
	router := chi.NewRouter()

	// Adding the routes essentially
	router.Get("/", GetPhrase)

	// Creating our route
	handlers.CreateRoute(router)
}

// The route callback, it will give us the content on `src/data/phrases.json`
func GetPhrase(response http.ResponseWriter, request *http.Request) {
	file, _ := ioutil.ReadFile("src/data/phrases.json")
	response.Write(file)
}
