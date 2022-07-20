package RouteDomain

import "net/http"

type Route struct {
	URI           string
	Method        string
	RouteFunction func(w http.ResponseWriter, r *http.Request)
	Authenticated bool
}
