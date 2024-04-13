package routes

import "net/http"

var loginRoute = DefaultRoute{
	URI: "/login",
	Method: http.MethodPost,
	Function: func(w http.ResponseWriter, r *http.Request) {},
	AuthRequired: false,
}