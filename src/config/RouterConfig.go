package config

import (
	"flightApi/src/controller"
	"github.com/gorilla/mux"
)

func RouterConfig(r *mux.Router) *mux.Router {
	routes := controller.UserControllers

	for _, route := range routes {
		r.HandleFunc(route.URI, route.RouteFunction).Methods(route.Method)
	}
	return r
}
