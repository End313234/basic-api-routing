package phrases

import (
	"io/ioutil"
	"net/http"

	"github.com/End313234/basic-api-routing/src/handlers"
)

func init() {
	// Creating the route
	route := handlers.Route{
		Method:  "GET",
		Pattern: "/phrases",
		Handler: GetAllPhrases,
	}
	route.Create()

	// Optional usage
	/*

		route := handlers.Route{}
		route.AddMethod("GET").AddPattern("/").AddHandler(GetPhrase).Create()

	*/
}

type GetAllPhrasesOutput = []string

// The route callback, it will give us the content on `src/data/phrases.json`
func GetAllPhrases(response http.ResponseWriter, request *http.Request) {
	file, _ := ioutil.ReadFile("src/data/phrases.json")
	response.Write(file)
}
