package utils

import (
	"encoding/json" // Importing encoding/json package for JSON parsing
	"io/ioutil"     // Importing ioutil package for reading request body
	"net/http"      // Importing net/http package for HTTP request handling
)

// ParseBody parses the JSON body of an HTTP request into the provided interface{}
func ParseBody(r *http.Request, x interface{}) {
	// Reading the request body
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		// Unmarshalling the JSON body into the provided interface{}
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return // Returning if there is an error in unmarshalling
		}
	}
}
