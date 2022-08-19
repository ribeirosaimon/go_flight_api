package routes

import (
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	URI           string
	Method        string
	Function      func(ctx *fiber.Ctx) error
	Authenticated bool
}

func RoutersConfig(app *fiber.App) *fiber.App {
	routers := UserRouters

	for _, route := range routers {
		app.Add(route.Method, route.URI, route.Function)
	}

	return app
}
