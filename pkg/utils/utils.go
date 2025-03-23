package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody reads the request body and unmarshals it into the given struct
func ParseBody(r *http.Request, x interface{}) {
	// Read the entire request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return // Exit if reading fails
	}

	// Unmarshal the JSON body into the provided struct
	if err := json.Unmarshal(body, x); err != nil {
		return // Exit if JSON parsing fails
	}
}
