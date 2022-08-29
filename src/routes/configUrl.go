package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/ribeirosaimon/go_flight_api/src/controllers/config"
)

func ConfigureApi(apiHandlers fiber.Router) {
	apiHandlers.Get("/swagger/*", swagger.HandlerDefault)
	apiHandlers.Get("/config", config.ApiUppController)
}
