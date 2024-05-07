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
		AuthRequired: true,
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
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/follow",
		Method:       http.MethodPost,
		Function:     controllers.FollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/unfollow",
		Method:       http.MethodPost,
		Function:     controllers.UnFollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/followers",
		Method:       http.MethodGet,
		Function:     controllers.SearchFollowers,
		AuthRequired: true,
	},
}
