package login

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/response"
	"github.com/ribeirosaimon/go_flight_api/src/security"
	"net/http"
	"strings"
)

const (
	_BODY_LOGIN_ERROR = "body needs username and password"
)

type controllerLogin struct {
	service serviceLogin
}

func ControllerLogin() controllerLogin {
	return controllerLogin{service: ServiceLogin()}
}

func (s controllerLogin) Login(c *fiber.Ctx) error {
	var account model.LoginDto
	if err := c.BodyParser(&account); err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: _BODY_LOGIN_ERROR})
	}

	accessToken, err := s.service.UserLogin(account)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(accessToken)
}

func (s controllerLogin) SignUp(c *fiber.Ctx) error {
	var account model.AccountDto
	if err := c.BodyParser(&account); err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	if strings.Contains(account.Username, " ") {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "your username contains space"})
	}

	if account.Username == "" {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "you need send username"})
	}
	encriptedPassword, err := security.EncriptyPassword(account.Password)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	account.Password = string(encriptedPassword)
	save, err := s.service.SaveOneAccount(account)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(save.SanitizerAccount())
}
