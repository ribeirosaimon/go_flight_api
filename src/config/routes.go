package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/routes"
)

type Route struct {
	URI           string
	Method        string
	Function      func(ctx *fiber.Ctx) error
	Authenticated bool
}

func routersConfig(app *fiber.App) *fiber.App {
	routers := routes.UserRouters

	for _, route := range routers {
		app.Add(route.Method, route.URI, route.Function)
	}

	return app
}
