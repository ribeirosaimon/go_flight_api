package repository

import (
	"context"
	"github.com/ribeirosaimon/go_flight_api/src/config"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const _FLIGHT_REPOSITORY = "flight"

type FlightRepositoryImpl struct {
	conn *mongo.Collection
}

func FlightRepository() FlightRepositoryImpl {
	return FlightRepositoryImpl{conn: config.GetMongoClient(_FLIGHT_REPOSITORY)}
}

func (mongo FlightRepositoryImpl) Save(ctx context.Context, flight model.Flight) (model.Flight, error) {
	one, err := mongo.conn.InsertOne(ctx, flight)
	if err != nil {
		return model.Flight{}, err
	}
	flight.ID = one.InsertedID.(primitive.ObjectID)
	return flight, nil
}

func (mongo FlightRepositoryImpl) FindById(ctx context.Context, id string) (model.Flight, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Flight{}, err
	}
	result := model.Flight{}
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}
	if err := mongo.conn.FindOne(ctx, filter).Decode(&result); err != nil {
		return result, err
	}
	return result, nil
}

func (mongo FlightRepositoryImpl) FindAll(ctx context.Context) ([]model.Flight, error) {
	var result []model.Flight
	filter := bson.D{{}}

	allFlights, err := mongo.conn.Find(ctx, filter)
	defer allFlights.Close(ctx)

	if err != nil {
		return nil, err
	}

	for allFlights.Next(ctx) {
		var flight model.Flight
		if err := allFlights.Decode(&flight); err != nil {
			return result, err
		}
		result = append(result, flight)
	}
	return result, nil
}

func (mongo FlightRepositoryImpl) DeleteById(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", objectId}}

	isDeleted, err := mongo.conn.DeleteOne(ctx, filter)

	if isDeleted.DeletedCount == 0 {
		return err
	}

	return nil
}

func (mongo FlightRepositoryImpl) FindMoreCheapFlight(ctx context.Context) (model.Flight, error) {
	filter := bson.D{}
	opts := options.FindOne().SetSort(bson.D{{"price", 1}})

	result := model.Flight{}
	if err := mongo.conn.FindOne(ctx, filter, opts).Decode(&result); err != nil {
		return result, err
	}
	return result, nil
}
