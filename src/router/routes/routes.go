package routes

import "net/http"

// DefaultRoute represent all API routes
type DefaultRoute struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}
