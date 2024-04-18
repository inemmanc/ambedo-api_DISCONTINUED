package routes

import (
	"ambedo-api/src/controllers"
	"net/http"
)

var loginRoute = DefaultRoute{
	URI:          "/login",
	Method:       http.MethodPost,
	Function:     controllers.Login,
	AuthRequired: false,
}
