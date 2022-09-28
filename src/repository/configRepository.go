package repository

import (
	"context"
	"github.com/ribeirosaimon/go_flight_api/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const _CONFIG_REPOSITORY = "configuration"

type ConfigRepositoryImpl struct {
	conn *mongo.Collection
}

type ConfigurationModel struct {
	Id        uint8 `bson:"_id" json:"id"`
	IsLoading bool  `bson:"isLoading" json:"isLoading"`
}

func ConfigRepository() ConfigRepositoryImpl {
	return ConfigRepositoryImpl{conn: config.GetMongoClient(_CONFIG_REPOSITORY)}
}

func (mongo ConfigRepositoryImpl) GetConfig(ctx context.Context) (ConfigurationModel, error) {
	result := ConfigurationModel{}
	filter := bson.D{primitive.E{Key: "_id", Value: 1}}

	err := mongo.conn.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return ConfigurationModel{}, err
	}
	return result, nil
}
