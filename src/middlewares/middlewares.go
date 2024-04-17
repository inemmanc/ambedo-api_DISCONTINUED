package middlewares

import (
	"fmt"
	"net/http"
)

// Authenticate checks whether the requesting user is authenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TEMP RESPONSE ---- REMOVE
		fmt.Println("VALIDATING ...")
		next(w, r)
	}
}