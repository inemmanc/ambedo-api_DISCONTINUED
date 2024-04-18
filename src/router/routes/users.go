package routes

import (
	"ambedo-api/src/controllers"
	"net/http"
)

var usersRoutes = []DefaultRoute{
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.FindUsers,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodGet,
		Function:     controllers.FindUser,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequired: false,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: false,
	},
}
