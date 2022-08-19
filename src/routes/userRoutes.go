package routes

import (
	"fmt"
	"github.com/ribeirosaimon/go_flight_api/src/controllers"
	"net/http"
)

var _USERROUTER = "/user"

var UserRouters = []Route{
	{
		URI:           fmt.Sprint(_USERROUTER + "/"),
		Method:        http.MethodGet,
		Function:      controllers.FindAllController,
		Authenticated: false,
	},
	{
		URI:           fmt.Sprint(_USERROUTER + "/:id"),
		Method:        http.MethodGet,
		Function:      controllers.FindOneUserController,
		Authenticated: false,
	},
	{
		URI:           fmt.Sprint(_USERROUTER + "/save"),
		Method:        http.MethodPost,
		Function:      controllers.SaveUserController,
		Authenticated: false,
	},
	{
		URI:           fmt.Sprint(_USERROUTER + "/:id"),
		Method:        http.MethodPut,
		Function:      controllers.UpdateUserController,
		Authenticated: false,
	},
}
