package routes

import (
	"fmt"
	"github.com/ribeirosaimon/go_flight_api/src/controllers/user"
	"net/http"
)

var _USERROUTER = "/user"

var UserRouters = []Route{
	{
		URI:           fmt.Sprint(_USERROUTER + "/"),
		Method:        http.MethodGet,
		Function:      user.FindAllController,
		Authenticated: true,
	},
	{
		URI:           fmt.Sprint(_USERROUTER + "/:id"),
		Method:        http.MethodGet,
		Function:      user.FindOneUserController,
		Authenticated: true,
	},
	{
		URI:           fmt.Sprint(_USERROUTER + "/save"),
		Method:        http.MethodPost,
		Function:      user.SaveUserController,
		Authenticated: false,
	},
	{
		URI:           fmt.Sprint(_USERROUTER + "/:id"),
		Method:        http.MethodPut,
		Function:      user.UpdateUserController,
		Authenticated: true,
	},

	{
		URI:           fmt.Sprint(_USERROUTER + "/:id"),
		Method:        http.MethodDelete,
		Function:      user.DeleteUserController,
		Authenticated: true,
	},
}
