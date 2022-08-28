package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/controllers/user"
)

func UserRoutes(apiHandlers fiber.Router) {
	apiHandlers = apiHandlers.Group("/user")

	apiHandlers.Get("", AdminPermission, user.FindAllController)
	apiHandlers.Get("/:id", UserPermission, user.FindOneUserController)
	apiHandlers.Put("/:id", UserPermission, user.UpdateUserController)
	apiHandlers.Delete("/:id", UserPermission, user.DeleteUserController)
}
