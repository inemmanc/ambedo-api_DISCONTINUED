package routes

import "net/http"

var usersRoutes = []DefaultRoute{
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     func(w http.ResponseWriter, r *http.Request) {},
		AuthRequired: false,
	},
	{
		URI:          "/user/{userID}",
		Method:       http.MethodGet,
		Function:     func(w http.ResponseWriter, r *http.Request) {},
		AuthRequired: false,
	},
	{
		URI:          "/user/{userID}",
		Method:       http.MethodPost,
		Function:     func(w http.ResponseWriter, r *http.Request) {},
		AuthRequired: false,
	},
	{
		URI:          "/user/{userID}",
		Method:       http.MethodPut,
		Function:     func(w http.ResponseWriter, r *http.Request) {},
		AuthRequired: false,
	},
	{
		URI:          "/user/{userID}",
		Method:       http.MethodDelete,
		Function:     func(w http.ResponseWriter, r *http.Request) {},
		AuthRequired: false,
	},
}
