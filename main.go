package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/routes"
)

func main() {
	app := fiber.New()

	routes.RoutersConfig(app)

	app.Listen(":3000")

}
