package flight

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/response"
	"net/http"
)

type flightController struct {
	flightService flightService
}

func Controller() flightController {
	return flightController{flightService: CreateFlightService()}
}

func (s flightController) FindAllController(c *fiber.Ctx) error {
	flights, err := s.flightService.findAllFlights()
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(flights)
}

func (s flightController) FindOneFlightController(c *fiber.Ctx) error {
	id := fmt.Sprint(c.Params("id"))
	flight, err := s.flightService.findById(id)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "user not found"})
	}
	return c.Status(http.StatusOK).JSON(flight)
}

func (s flightController) SaveFlightController(c *fiber.Ctx) error {
	var flight model.Flight
	if err := c.BodyParser(&flight); err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	saveFlight, err := s.flightService.saveFlight(flight)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(saveFlight)
}

func (s flightController) DeleteFlightController(c *fiber.Ctx) error {
	id := fmt.Sprint(c.Params("id"))
	err := s.flightService.deleteById(id)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.SendStatus(http.StatusOK)
}

func (s flightController) CheapFlight(c *fiber.Ctx) error {
	cheapFlight, err := s.flightService.cheapFlight()
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(cheapFlight)
}
