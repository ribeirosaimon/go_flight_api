package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/model"
)

func UserPermission(ctx *fiber.Ctx) error {
	return Authorization(ctx, []string{model.USER})
}
func AdminPermission(ctx *fiber.Ctx) error {
	return Authorization(ctx, []string{model.ADMIN})
}
