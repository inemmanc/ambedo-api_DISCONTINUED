package router

import (
	"ambedo-api/src/router/routes"

	"github.com/gorilla/mux"
)

// returns a router with all configured routes
func GenerateRouter() *mux.Router {
	newRouterInstance := mux.NewRouter()

	return routes.Configure(newRouterInstance)
}
