package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/repository"
)

func CreateFlight(c *fiber.Ctx) error {
	flights, err := repository.ReadFlights()
	if err != nil {
		return err
	}
	return c.JSON(flights)
}
