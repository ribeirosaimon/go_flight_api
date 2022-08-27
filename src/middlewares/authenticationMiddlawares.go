package middlewares

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/response"
	"github.com/ribeirosaimon/go_flight_api/src/security"
	"net/http"
	"strings"
)

func Authentication(c *fiber.Ctx) error {
	loggedUser, err := getUser(c)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "token was expired"})
	}
	c.Locals("loggedUser", loggedUser)
	return c.Next()
}

func Authorization(c *fiber.Ctx, roles []string) error {
	loggedUser, err := getUser(c)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "token was expired"})
	}
	authorization := contains(roles, loggedUser.Roles)
	fmt.Println(authorization)
	fmt.Println(roles)
	if !authorization {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "you not have permission"})
	}
	return c.Next()
}

func contains(loggedUserRole, routeRole []string) bool {
	for _, userRole := range loggedUserRole {

		for _, role := range routeRole {
			if userRole == role {
				return true
			}
		}
	}
	return false
}

func getUser(c *fiber.Ctx) (model.LoggedUser, error) {
	var token string
	headerAuthorization := c.Get("Authorization")

	if len(strings.Split(headerAuthorization, " ")) == 2 {
		token = strings.Split(headerAuthorization, " ")[1]
	} else {
		return model.LoggedUser{}, errors.New("you need a access token")
	}

	return security.ValidationToken(token)
}
