package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/controllers/login"
	"github.com/ribeirosaimon/go_flight_api/src/middlewares"
	"github.com/ribeirosaimon/go_flight_api/src/model"
)

type Route struct {
	URI             string
	Method          string
	Function        func(ctx *fiber.Ctx) error
	Authenticated   bool
	RoutePermission []string
}

func AddApiRoutes(app *fiber.App) {
	apiHandlers := app.Group("/api")

	LoginRouter(apiHandlers)
	UserRoutes(apiHandlers)
	SwaggerRoutes(apiHandlers)

}

func LoginRouter(apiHandlers fiber.Router) {
	apiHandlers.Post("/login", login.ControllerLogin)
	apiHandlers.Post("/signup", login.SignUp)
}

func UserPermission(ctx *fiber.Ctx) error {
	return middlewares.Authorization(ctx, []string{model.USER})
}
func AdminPermission(ctx *fiber.Ctx) error {
	return middlewares.Authorization(ctx, []string{model.ADMIN})
}
