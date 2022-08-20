package middlewares

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/response"
	"github.com/ribeirosaimon/go_flight_api/src/security"
	"net/http"
	"strings"
)

const (
	_BODY_LOGIN_ERROR = "you need be authenticated"
)

func Authentication(c *fiber.Ctx) error {
	var token string
	headerAuthorization := c.Get("Authorization")

	if len(strings.Split(headerAuthorization, " ")) == 2 {
		token = strings.Split(headerAuthorization, " ")[1]
	} else {
		return errors.New("you need a access token")
	}

	loggedUser, err := security.ValidationToken(token)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "token was expired"})
	}
	c.Locals("loggedUser", loggedUser)
	return c.Next()

}
