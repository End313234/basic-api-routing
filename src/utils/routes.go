package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Reading the request body (so I don't need to write this whole code for each route)
func ReadRequestBody(body io.ReadCloser, response http.ResponseWriter, v interface{}) error {
	stringifiedBody, _ := ioutil.ReadAll(body)

	err := json.Unmarshal(stringifiedBody, v)
	if err != nil {
		ThrowError(
			Error{
				StatusCode: 400,
				Message:    "error while reading request body",
			},
			response,
		)
		return err
	}

	return nil
}

// Usage example:
/*

type Input struct {
	parameter string
}

func RouteCallback(response http.ResponseWriter, request *http.Request) {
	input := Input{}
	err := utils.ReadRequestBody(request.Body, response, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(input)
}

*/
