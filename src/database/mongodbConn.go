package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	StringConn = ""
	database   = "FlightApi"
)

func GetCollection(collection string) *mongo.Collection {
	var err error
	//url := os.Getenv("MONGO_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/FlightApi"))
	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err.Error())
	}
	return client.Database(database).Collection(collection)
}
