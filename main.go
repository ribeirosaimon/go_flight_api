package main

import (
	"context"
	"fmt"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/repository"
)

func main() {

	var user = model.Account{
		Name:     "teste",
		LastName: "teste",
		Password: "pass",
	}

	mongoRepository := repository.NewMongoRepository()
	save, err := mongoRepository.Account.Save(context.Background(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println(save)

	//app := fiber.New()
	//app.Use(cors.New(cors.Config{
	//	AllowOrigins: "*",
	//	AllowHeaders: "Origin, Content-Type, Accept",
	//}))
	//
	//routes.AddApiRoutes(app)
	//
	//err := app.Listen(":3000")
	//if err != nil {
	//	return
	//}

}
