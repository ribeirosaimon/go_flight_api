package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/middlewares"
)

func UserRoutes(apiHandlers fiber.Router) {
	apiHandlers = apiHandlers.Group("/user")

	apiHandlers.Get("", middlewares.AdminPermission, UserControler().FindAllController)
	apiHandlers.Get("/:id", middlewares.UserPermission, UserControler().FindOneUserController)
	apiHandlers.Put("/:id", middlewares.UserPermission, UserControler().UpdateUserController)
	apiHandlers.Delete("/:id", middlewares.UserPermission, UserControler().DeleteUserController)
	apiHandlers.Post("/:id/promoted", middlewares.UserPermission, UserControler().PromotedToAdmin)
}
