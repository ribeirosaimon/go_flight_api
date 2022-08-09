package routers

import "github.com/gofiber/fiber/v2"

type Router struct {
	URI                string
	Method             string
	Function           func(*fiber.Ctx) error
	NeedAuthentication bool
}

func ConfigRouter(r *fiber.App) *fiber.App {
	routes := FlighRouter

	for _, route := range routes {
		r.Add(route.Method, route.URI, route.Function)
	}

	return r
}
