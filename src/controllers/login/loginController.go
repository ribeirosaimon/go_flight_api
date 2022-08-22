package login

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/response"
	"github.com/ribeirosaimon/go_flight_api/src/services"
	"net/http"
)

const (
	_BODY_LOGIN_ERROR = "body needs username and password"
)

func ControllerLogin(c *fiber.Ctx) error {
	var user model.LoginDto
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: _BODY_LOGIN_ERROR})
	}
	accessToken, err := services.UserLogin(user)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(accessToken)
}
