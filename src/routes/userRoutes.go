package routes

import (
	"fmt"
	"github.com/ribeirosaimon/go_flight_api/src/config"
	"github.com/ribeirosaimon/go_flight_api/src/controllers"
	"net/http"
)

var _USERROUTER = "/user/"

var UserRouters = []config.Route{
	{
		URI:           fmt.Sprint(_USERROUTER + "/save"),
		Method:        http.MethodPost,
		Function:      controllers.SaveUser,
		Authenticated: false,
	},
}
