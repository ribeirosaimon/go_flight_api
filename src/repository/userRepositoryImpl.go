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
)

const _USER_COLLECTION = "account"

type userRepository struct {
	conn *mongo.Collection
}

func newUserRepository() *userRepository {
	return &userRepository{conn: config.GetMongoClient(_USER_COLLECTION)}
}

func (mongo userRepository) Save(ctx context.Context, account model.Account) (model.Account, error) {
	err := mongo.isDuplicateUser(ctx, account.Username)
	if err != nil {
		return model.Account{}, err
	}

	one, err := mongo.conn.InsertOne(ctx, account)

	account.ID = one.InsertedID.(primitive.ObjectID)
	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}

func (mongo userRepository) FindById(ctx context.Context, ID string) (model.Account, error) {
	result := model.Account{}
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return result, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

	if err := mongo.conn.FindOne(ctx, filter).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func (mongo userRepository) FindByUsername(ctx context.Context, username string) (model.Account, error) {
	result := model.Account{}
	filter := bson.D{primitive.E{Key: "username", Value: username}}

	if err := mongo.conn.FindOne(ctx, filter).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func (mongo userRepository) FindAll(ctx context.Context) ([]model.Account, error) {
	var allResults []model.Account

	find, err := mongo.conn.Find(ctx, bson.D{{}})
	defer find.Close(ctx)

	if err != nil {
		return nil, err
	}

	for find.Next(ctx) {
		var account model.Account
		if err := find.Decode(&account); err != nil {
			return allResults, err
		}
		allResults = append(allResults, account)
	}

	return allResults, nil
}

func (mongo userRepository) Update(ctx context.Context, ID string, account model.AccountDto) (model.Account, error) {
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return model.Account{}, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}
	itensForUpdate := bson.D{}

	err = mongo.isDuplicateUser(ctx, account.Username)
	if err != nil {
		return model.Account{}, err
	}

	if account.Name != "" {
		itensForUpdate = append(itensForUpdate, primitive.E{Key: "name", Value: account.Name})
	}
	if account.Username != "" {
		itensForUpdate = append(itensForUpdate, primitive.E{Key: "username", Value: account.Username})
	}
	if account.LastName != "" {
		itensForUpdate = append(itensForUpdate, primitive.E{Key: "lastName", Value: account.LastName})
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

	result := mongo.conn.FindOneAndUpdate(ctx, filter, bsonUpdate, &opt)

	if result.Err() != nil {
		fmt.Println(result)
		return model.Account{}, err
	}
	var updatedAccount model.Account
	decodeErr := result.Decode(&updatedAccount)
	return updatedAccount, decodeErr
}

func (mongo userRepository) Delete(ctx context.Context, ID string) error {
	objectId, err := primitive.ObjectIDFromHex(ID)
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

	_, err = mongo.conn.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (mongo userRepository) isDuplicateUser(ctx context.Context, username string) error {
	duplicatedUser := bson.D{primitive.E{Key: "username", Value: username}}
	documents, err := mongo.conn.CountDocuments(ctx, duplicatedUser)
	if err != nil {
		return err
	}
	if documents != 0 {
		return errors.New("user already exists")
	}
	return nil
}
