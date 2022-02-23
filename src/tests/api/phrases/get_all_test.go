package phrases

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

const BASE_URL = "http://localhost:3000"

func TestGetAllPhrases(test *testing.T) {
	response, err := http.Get(BASE_URL + "/phrases")
	if err != nil {
		test.Log("could not make request (are you sure the api is online?)")
		test.Fail()
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	phrases := []string{}
	if err := json.Unmarshal(body, &phrases); err != nil {
		test.Log("wrong response body type (must be array of strings)")
		test.Fail()
	}

	file, _ := ioutil.ReadFile("../../../data/phrases.json")
	parsedFile := []string{}

	if json.Unmarshal(file, &parsedFile); !reflect.DeepEqual(phrases, parsedFile) {
		test.Log("given phrases are incorrect")
		test.Fail()
	}
}
