package repository

import (
	"context"
	"errors"
	"github.com/ribeirosaimon/go_flight_api/src/config"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	_ACCOUNTCONNECTION = "account"
	_TIMEOUTCONTEXT    = 2
)

func FindAll() ([]model.Account, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	filter := bson.D{{}}
	var allResults []model.Account
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

		allResults = append(allResults, account.SanitizerAccount())
	}
	cur.Close(ctx)
	if len(allResults) == 0 {
		return allResults, mongo.ErrNoDocuments
	}
	return allResults, nil
}

func FindById(id string) (model.Account, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	result := model.Account{}
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

	client, err := config.GetMongoClient()
	if err != nil {
		return result, err
	}
	collection := client.Database(config.DB).Collection(_ACCOUNTCONNECTION)

	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return result, err
	}
	return result.SanitizerAccount(), err
}

//Save my Account in Repository and return this Account
func Save(account model.Account) (model.Account, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	var newAccount = model.Account{}
	client, err := config.GetMongoClient()
	if err != nil {
		return newAccount, err
	}

	collection := client.Database(config.DB).Collection(_ACCOUNTCONNECTION)
	resp, err := collection.InsertOne(ctx, account)
	if err != nil {
		return newAccount, err
	}

	newAccount.ID = resp.InsertedID.(primitive.ObjectID)
	newAccount.Name = account.Name
	newAccount.LastName = account.LastName

	return newAccount, nil
}

func Update(id string, account model.AccountDto) (model.Account, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	client, err := config.GetMongoClient()
	if err != nil {
		return model.Account{}, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}
	updater := bson.D{primitive.E{Key: "$set",
		Value: bson.D{
			primitive.E{Key: "name", Value: account.Name},
			primitive.E{Key: "lastName", Value: account.LastName},
			primitive.E{Key: "password", Value: account.Password},
			primitive.E{Key: "updatedAt", Value: time.Now()},
		},
	}}

	collection := client.Database(config.DB).Collection(_ACCOUNTCONNECTION)

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := collection.FindOneAndUpdate(ctx, filter, updater, &opt)
	if result.Err() != nil {
		return model.Account{}, errors.New(err.Error())
	}
	var updatedAccount model.Account
	decodeErr := result.Decode(&updatedAccount)
	return updatedAccount.SanitizerAccount(), decodeErr
}

func Delete(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

	client, err := config.GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database(config.DB).Collection(_ACCOUNTCONNECTION)
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
