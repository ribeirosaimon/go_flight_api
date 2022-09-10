package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ribeirosaimon/go_flight_api/src/routes"
)

func main() {

	//var user = model.AccountDto{
	//	Name:     "teste",
	//	LastName: "teste",
	//	Password: "pass",
	//}
	//
	//service := services.UserService()
	//account, err := service.SaveOneAccount(user)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(account)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.AddApiRoutes(app)

	err := app.Listen(":4000")
	if err != nil {
		return
	}

}
