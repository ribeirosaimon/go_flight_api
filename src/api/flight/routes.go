package flight

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/middlewares"
)

func FlightRoutes(apiHandlers fiber.Router) {
	apiHandlers = apiHandlers.Group("/flight")

	apiHandlers.Get("", middlewares.AdminPermission, Controller().FindAllController)
	apiHandlers.Post("", middlewares.UserPermission, Controller().SaveFlightController)
	apiHandlers.Get("/cheap", middlewares.UserPermission, Controller().CheapFlight)
	apiHandlers.Post("/search", middlewares.UserPermission, Controller().SearchFlight)
	apiHandlers.Get("/last-flight", middlewares.UserPermission, Controller().GetLastFlight)
	apiHandlers.Get("/:id", middlewares.UserPermission, Controller().FindOneFlightController)
	apiHandlers.Delete("/:id", middlewares.AdminPermission, Controller().DeleteFlightController)
}
