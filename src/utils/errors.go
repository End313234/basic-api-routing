package utils

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// Throwing an error
func ThrowError(err Error, response http.ResponseWriter) {
	parsedError, _ := json.Marshal(err)
	response.Write(parsedError)
}
