package user

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/response"
	"github.com/ribeirosaimon/go_flight_api/src/services"
	"net/http"
)

const (
	_ERRO_IN_BODY   = "error in your body"
	_NOT_FOUND_USER = "user not found"
)

func FindAllController(c *fiber.Ctx) error {
	user, err := services.FindAllUserService()
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(user)
}

func FindOneUserController(c *fiber.Ctx) error {
	id := fmt.Sprint(c.Params("id"))
	_, err := validateLoggedUser(c, id)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	user, err := services.FindOneUserService(id)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: _NOT_FOUND_USER})
	}
	return c.Status(http.StatusOK).JSON(user.SanitizerAccount())
}

func UpdateUserController(c *fiber.Ctx) error {
	id := fmt.Sprint(c.Params("id"))
	_, err := validateLoggedUser(c, id)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	var user model.AccountDto
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: _ERRO_IN_BODY})
	}
	updatedUser, err := services.UpdateUserService(id, user)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(updatedUser.SanitizerAccount())
}

func DeleteUserController(c *fiber.Ctx) error {
	id := fmt.Sprint(c.Params("id"))
	_, err := validateLoggedUser(c, id)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	err = services.DeleteUserService(id)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: _NOT_FOUND_USER})
	}
	return c.SendStatus(http.StatusOK)
}

func validateLoggedUser(ctx *fiber.Ctx, searchId string) (model.LoggedUser, error) {
	val := ctx.Locals("loggedUser")
	loggedUser := val.(model.LoggedUser)

	if loggedUser.UserId != searchId {
		return model.LoggedUser{}, errors.New("you not have permission")
	}
	return loggedUser, nil
}
