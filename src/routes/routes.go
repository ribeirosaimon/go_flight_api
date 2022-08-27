package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/middlewares"
)

type Route struct {
	URI             string
	Method          string
	Function        func(ctx *fiber.Ctx) error
	Authenticated   bool
	RoutePermission []string
}

func RoutersConfig(app *fiber.App) *fiber.App {
	routers := UserRouters
	routers = append(routers, LoginRoute)
	for _, route := range routers {
		apiHandlers := app.Group("/api")
		)
		//if route.Authenticated {
		//	apiHandlers.Add(
		//		route.Method,
		//		route.URI,
		//		middlewares.Authentication,
		//		func(ctx *fiber.Ctx) error {
		//			return middlewares.Authorization(ctx, route.RoutePermission)
		//		},
		//		route.Function)
		//} else {
		//	apiHandlers.Add(route.Method, route.URI, route.Function)
		//}
	}
	return app
}
