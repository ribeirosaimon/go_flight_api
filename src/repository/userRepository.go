package repository

import (
	"context"
	"errors"
	"fmt"
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
)

type DatabaseContext struct {
	context context.Context
}

func FindAll() ([]model.Account, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	filter := bson.D{{}}
	var allResults []model.Account
	client := config.GetMongoClient()
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

	client := config.GetMongoClient()
	collection := client.Database(config.DB).Collection(_ACCOUNTCONNECTION)
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return result, err
	}
	return result.SanitizerAccount(), err
}

func FindUserByUsername(username string) (model.Account, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	result := model.Account{}
	filter := bson.D{primitive.E{Key: "username", Value: username}}

	client := config.GetMongoClient()
	collection := client.Database(config.DB).Collection(_ACCOUNTCONNECTION)
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, err
}

//Save my Account in Repository and return this Account
func Save(account model.Account) (model.Account, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	var newAccount = model.Account{}
	client := config.GetMongoClient()

	collection := client.Database(config.DB).Collection(_ACCOUNTCONNECTION)

	err := duplicatedUser(ctx, account.Username, collection)
	if err != nil {
		return model.Account{}, err
	}

	resp, err := collection.InsertOne(ctx, account)
	if err != nil {
		return model.Account{}, errors.New(fmt.Sprintf("exists user with your username :%s", account.Username))
	}

	newAccount.ID = resp.InsertedID.(primitive.ObjectID)
	newAccount.Name = account.Name
	newAccount.LastName = account.LastName
	newAccount.Username = account.Username
	newAccount.Roles = account.Roles

	return newAccount, nil
}

func Update(id string, account model.AccountDto) (model.Account, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	client := config.GetMongoClient()

	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}
	itensForUpdate := bson.D{}

	if account.Name != "" {
		itensForUpdate = append(itensForUpdate, primitive.E{Key: "name", Value: account.Name})
	}
	if account.Username != "" {
		itensForUpdate = append(itensForUpdate, primitive.E{Key: "username", Value: account.Username})
	}
	if account.LastName != "" {
		itensForUpdate = append(itensForUpdate, primitive.E{Key: "lastName", Value: account.LastName})
	}
	itensForUpdate = append(itensForUpdate, primitive.E{Key: "updatedAt", Value: time.Now()})

	collection := client.Database(config.DB).Collection(_ACCOUNTCONNECTION)
	err = duplicatedUser(ctx, account.Username, collection)

	if err != nil {
		return model.Account{}, err
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	bsonUpdate := bson.D{primitive.E{
		Key: "$set", Value: itensForUpdate,
	}}

	result := collection.FindOneAndUpdate(ctx, filter, bsonUpdate, &opt)
	if result.Err() != nil {
		fmt.Println(result)
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

	client := config.GetMongoClient()

	collection := client.Database(config.DB).Collection(_ACCOUNTCONNECTION)
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
func duplicatedUser(ctx context.Context, username string, collection *mongo.Collection) error {
	duplicatedUser := bson.D{primitive.E{Key: "username", Value: username}}
	documents, err := collection.CountDocuments(ctx, duplicatedUser)
	if err != nil {
		return err
	}
	if documents != 0 {
		return errors.New(fmt.Sprintf("exists user with your username :%s", username))
	}
	return nil
}
