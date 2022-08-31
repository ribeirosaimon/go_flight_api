package main

import (
	"context"
	"github.com/ribeirosaimon/go_flight_api/src/config"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/repository"
	"log"
)

func main() {
	repo := repository.New(repository.Options{WriteMongo: config.GetMongoClient(), DatabaseName: config.DB, Collection: "account"})
	var user = model.Account{
		Name:     "teste",
		LastName: "teste",
		Password: "pass",
	}

	err := repo.User.Create(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
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
