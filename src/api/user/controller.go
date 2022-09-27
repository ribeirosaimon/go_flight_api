package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/api/login"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/response"
	"net/http"
)

type userController struct {
	userService userService
}

func UserControler() userController {
	return userController{userService: UserService()}
}

func (s userController) FindAllController(c *fiber.Ctx) error {

	user, err := s.userService.FindAllUserService()
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(user)
}

func (s userController) FindOneUserController(c *fiber.Ctx) error {
	id := fmt.Sprint(c.Params("id"))
	if login.ValidateLoggedUser(c).UserId != id {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "You not have permission"})
	}
	user, err := s.userService.FindOneUserService(id)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(user.SanitizerAccount())
}

func (s userController) whoIsMe(c *fiber.Ctx) error {
	loggedUser := login.WhoIsMe(c)
	return c.Status(http.StatusOK).JSON(loggedUser)
}

func (s userController) UpdateUserController(c *fiber.Ctx) error {
	id := fmt.Sprint(c.Params("id"))
	if login.ValidateLoggedUser(c).UserId != id {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "You not have permission"})
	}
	var user model.AccountDto
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	updatedUser, err := s.userService.UpdateUserService(id, user)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(updatedUser.SanitizerAccount())
}

func (s userController) DeleteUserController(c *fiber.Ctx) error {
	id := fmt.Sprint(c.Params("id"))
	if login.ValidateLoggedUser(c).UserId != id {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "You not have permission"})
	}
	err := s.userService.DeleteUserService(id)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.SendStatus(http.StatusOK)
}

func (s userController) PromotedToAdmin(c *fiber.Ctx) error {
	id := fmt.Sprint(c.Params("id"))
	loggedUser := login.ValidateLoggedUser(c)
	if err := s.userService.promotedToAdmin(loggedUser, id); err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.SendStatus(http.StatusOK)
}

func (s userController) verifyPassword(c *fiber.Ctx) error {
	var user model.Account
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	loggedUser := login.ValidateLoggedUser(c)
	_, err := s.userService.verifyPassword(loggedUser, user)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: err.Error()})
	}
	return c.SendStatus(http.StatusOK)
}
