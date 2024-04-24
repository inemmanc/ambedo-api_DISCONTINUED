package routes

import (
	"ambedo-api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// struct represents the Default API route
type DefaultRoute struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// takes an empty Router and puts all new routes on it
func Configure(r *mux.Router) *mux.Router {
	var GeneralRoutes []DefaultRoute
	GeneralRoutes = append(GeneralRoutes, usersRoutes...)
	GeneralRoutes = append(GeneralRoutes, loginRoute)

	for _, route := range GeneralRoutes {

		if route.AuthRequired {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}

	}

	return r
}
