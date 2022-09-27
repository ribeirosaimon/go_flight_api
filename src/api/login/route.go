package login

import "github.com/gofiber/fiber/v2"

func LoginRouter(apiHandlers fiber.Router) {
	apiHandlers.Post("/login", ControllerLogin().Login)
	apiHandlers.Post("/signup", ControllerLogin().SignUp)

}
