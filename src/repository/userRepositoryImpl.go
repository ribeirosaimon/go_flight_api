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
)

const _USER_COLLECTION = "account"

type AccountRepositoryImpl struct {
	conn *mongo.Collection
}

func UserRepository() AccountRepositoryImpl {
	return AccountRepositoryImpl{conn: config.GetMongoClient(_USER_COLLECTION)}
}

func (mongo AccountRepositoryImpl) Save(ctx context.Context, account model.Account) (model.Account, error) {
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

func (mongo AccountRepositoryImpl) FindById(ctx context.Context, ID string) (model.Account, error) {
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

func (mongo AccountRepositoryImpl) FindByUsername(ctx context.Context, username string) (model.Account, error) {
	result := model.Account{}
	filter := bson.D{primitive.E{Key: "username", Value: username}}

	if err := mongo.conn.FindOne(ctx, filter).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func (mongo AccountRepositoryImpl) FindAll(ctx context.Context) ([]model.Account, error) {
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
		allResults = append(allResults, account.SanitizerAccount())
	}

	return allResults, nil
}

func (mongo AccountRepositoryImpl) Update(ctx context.Context, account model.Account) (model.Account, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: account.ID}}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	bsonUpdate := bson.D{primitive.E{
		Key: "$set", Value: account,
	}}

	result := mongo.conn.FindOneAndUpdate(ctx, filter, bsonUpdate, &opt)

	if result.Err() != nil {
		return model.Account{}, result.Err()
	}
	var updatedAccount model.Account
	decodeErr := result.Decode(&updatedAccount)
	return updatedAccount, decodeErr
}

func (mongo AccountRepositoryImpl) Delete(ctx context.Context, ID string) error {
	objectId, err := primitive.ObjectIDFromHex(ID)
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

	_, err = mongo.conn.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (mongo AccountRepositoryImpl) isDuplicateUser(ctx context.Context, username string) error {
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
