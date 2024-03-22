package router

import "github.com/gorilla/mux"

// GenerateRouter returns a router with all configured routes
func GenerateRouter() *mux.Router {
	return mux.NewRouter()
}
