package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/api/flight"
	"github.com/ribeirosaimon/go_flight_api/src/api/login"
	"github.com/ribeirosaimon/go_flight_api/src/api/user"
)

func AddApiRoutes(app *fiber.App) {
	apiHandlers := app.Group("/api")

	login.LoginRouter(apiHandlers)
	user.UserRoutes(apiHandlers)
	flight.FlightRoutes(apiHandlers)

	ConfigureApi(apiHandlers)
}
