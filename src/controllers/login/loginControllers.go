package login

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/response"
	"github.com/ribeirosaimon/go_flight_api/src/security"
	"github.com/ribeirosaimon/go_flight_api/src/services"
	"net/http"
	"strings"
)

const (
	_BODY_LOGIN_ERROR = "body needs username and password"
)

func ControllerLogin(c *fiber.Ctx) error {
	var user model.LoginDto
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: _BODY_LOGIN_ERROR})
	}
	accessToken, err := services.UserService().UserLogin(user)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(accessToken)
}

func SignUp(c *fiber.Ctx) error {
	var user model.AccountDto
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	if strings.Contains(user.Username, " ") {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "your username contains space"})
	}

	if user.Username == "" {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "you need send username"})
	}
	encriptedPassword, err := security.EncriptyPassword(user.Password)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	user.Password = string(encriptedPassword)
	save, err := services.UserService().SaveOneAccount(user)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(save)
}
