package config

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/repository"
	"github.com/ribeirosaimon/go_flight_api/src/response"
	"net/http"
	"time"
)

func ApiUppController(ctx *fiber.Ctx) error {
	myContext, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	config, err := repository.ConfigRepository().GetConfig(myContext)
	if err != nil {
		return ctx.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "Problem with get configuration"})
	}
	configApi := model.ConfigApiModel{
		Read:        true,
		CurrentTime: time.Now(),
		IsLoading:   config.IsLoading,
	}
	return ctx.Status(http.StatusOK).JSON(configApi)
}
