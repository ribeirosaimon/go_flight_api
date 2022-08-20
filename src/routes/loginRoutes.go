package routes

import (
	"github.com/ribeirosaimon/go_flight_api/src/controllers/login"
	"net/http"
)

var LoginRoute = Route{
	URI:           "/login",
	Method:        http.MethodGet,
	Function:      login.ControllerLogin,
	Authenticated: false,
}
