package controller

import (
	"flightApi/src/router/RouteDomain"
	"flightApi/src/services"
	"net/http"
)

var UserControllers = []RouteDomain.Route{
	{
		URI:           "/users",
		Method:        http.MethodGet,
		RouteFunction: services.CreateUser,
		Authenticated: false,
	},
}
