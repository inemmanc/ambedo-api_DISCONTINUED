package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns a json response to the request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
