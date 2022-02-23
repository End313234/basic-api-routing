package main

import (
	"fmt"
	"net/http"

	_ "github.com/End313234/basic-api-routing/src/api/phrases"
	"github.com/End313234/basic-api-routing/src/handlers"
	"github.com/End313234/basic-api-routing/src/utils"
	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	router.NotFound(func(response http.ResponseWriter, request *http.Request) {
		utils.ThrowError(
			utils.Error{
				StatusCode: 404,
				Message:    "Route not found",
			},
			response,
		)
	})
	router.MethodNotAllowed(func(response http.ResponseWriter, request *http.Request) {
		utils.ThrowError(
			utils.Error{
				StatusCode: 405,
				Message:    "Method not allowed",
			},
			response,
		)
	})

	routes := handlers.GetRouter()
	router.Mount("/", routes)

	fmt.Println("API ready!")
	http.ListenAndServe(":3000", router)
}
