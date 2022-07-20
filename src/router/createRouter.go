package router

import (
	"flightApi/src/config"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	return config.RouterConfig(mux.NewRouter())
}
