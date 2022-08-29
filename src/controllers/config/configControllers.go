package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"net/http"
	"time"
)

func ApiUppController(ctx *fiber.Ctx) error {
	configApi := model.ConfigApiModel{
		Read:        true,
		CurrentTime: time.Now(),
	}
	return ctx.Status(http.StatusOK).JSON(configApi)
}
