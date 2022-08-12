package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/config"
	"github.com/ribeirosaimon/go_flight_api/src/routers"
)

func main() {
	config.LoadConfig()
	app := fiber.New()
	routers.ConfigRouter(app)
	app.Listen(fmt.Sprintf(":%d", config.Port))
}
