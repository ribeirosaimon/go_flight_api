package repository

import (
	"context"
	"github.com/ribeirosaimon/go_flight_api/src/config"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	_ACCOUNTCONNECTION = "account"
	_TIMEOUTCONTEXT    = 2
)

//Todo Change all to interface{}

func FindAll() ([]model.Account, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	filter := bson.D{{}}
	allResults := []model.Account{}
	client, err := config.GetMongoClient()
	if err != nil {
		return allResults, err
	}
	collection := client.Database(config.DB).Collection(_ACCOUNTCONNECTION)

	cur, findError := collection.Find(ctx, filter)
	if findError != nil {
		return allResults, findError
	}

	for cur.Next(ctx) {
		account := model.Account{}
		if err := cur.Decode(&account); err != nil {
			return allResults, err
		}
		allResults = append(allResults, account)
	}
	cur.Close(ctx)
	if len(allResults) == 0 {
		return allResults, mongo.ErrNoDocuments
	}
	return allResults, nil
}

func FindById(id string) (model.Account, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	result := model.Account{}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	client, err := config.GetMongoClient()
	if err != nil {
		return result, err
	}
	collection := client.Database(config.DB).Collection(_ACCOUNTCONNECTION)

	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return result, err
	}
	return result, err
}

func Save(account model.Account) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	client, err := config.GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database(config.DB).Collection(_ACCOUNTCONNECTION)
	_, err = collection.InsertOne(ctx, account)
	if err != nil {
		return err
	}
	return nil
}

func Update(id string, account model.Account) (model.Account, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	account = model.Account{}

	client, err := config.GetMongoClient()
	if err != nil {
		return account, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	//https://levelup.gitconnected.com/working-with-mongodb-using-golang-754ead0c10c
	updater := bson.D{primitive.E{Key: "$set",
		Value: bson.D{primitive.E{Key: "name", Value: account.Name}}}}

}
