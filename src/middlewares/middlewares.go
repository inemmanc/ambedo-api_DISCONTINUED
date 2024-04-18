package middlewares

import (
	"fmt"
	"net/http"
)

// Logger logs all the request information (Method, URI, Host)
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\nLOG: %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authenticate checks whether the requesting user is authenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TEMP RESPONSE ---- REMOVE
		fmt.Println(" VALIDATING ...")
		next(w, r)
	}
}
