package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// DefaultRoute struct represents the Default API route
type DefaultRoute struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// Configure takes an empty Router and puts all new routes on it
func Configure(r *mux.Router) *mux.Router {
	var GeneralRoutes []DefaultRoute
	GeneralRoutes = append(GeneralRoutes, usersRoutes...)

	for _, route := range GeneralRoutes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
