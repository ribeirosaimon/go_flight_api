package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const (
	CONNSTR    = "MONGOCONN"
	DB         = "MONGODB"
	COLLECTION = "MONGOCOLLECTION"
)

type MongoConnect struct {
	Client *mongo.Client
}

func NewMongoConnect() (*MongoConnect, error) {
	M := MongoConnect{}
	cliops := options.Client().ApplyURI(os.Getenv(CONNSTR))
	client, err := mongo.Connect(context.TODO(), cliops)
	M.Client = client
	return &M, err
}

func (M *MongoConnect) DisconnectMongoClient() error {
	err := M.Client.Disconnect(context.TODO())
	if err != nil {
		return err
	}
	return nil
}
