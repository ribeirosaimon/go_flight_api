package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

const (
	CONNECTIONSTRING = "mongodb://localhost:27017"
	DB               = "flight_api_v1"
)

//GetMongoClient - Return mongodb connection to work with
func GetMongoClient() *mongo.Client {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			panic(err)
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			panic(err)
		}
		clientInstance = client
	})
	return clientInstance
}
