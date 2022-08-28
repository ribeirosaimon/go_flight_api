package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerRoutes(apiHandlers fiber.Router) {
	apiHandlers.Get("/swagger/*", swagger.HandlerDefault)
}
