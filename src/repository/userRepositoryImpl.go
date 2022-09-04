package repository

import (
	"context"
	"github.com/ribeirosaimon/go_flight_api/src/config"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const _USER_COLLECTION = "account"

type userRepository struct {
	conn *mongo.Collection
}

func newUserRepository() *userRepository {
	return &userRepository{conn: config.GetMongoClient(_USER_COLLECTION)}
}

func (mongo userRepository) Save(ctx context.Context, account model.Account) (model.Account, error) {
	one, err := mongo.conn.InsertOne(ctx, account)

	account.ID = one.InsertedID.(primitive.ObjectID)
	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}

func (mongo userRepository) FindById(ctx context.Context, ID string) (model.Account, error) {
	result := model.Account{}
	filter := bson.D{primitive.E{Key: "_id", Value: ID}}

	if err := mongo.conn.FindOne(ctx, filter).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func (mongo userRepository) FindAll(ctx context.Context) ([]model.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (mongo userRepository) Update(ctx context.Context, ID string, account model.Account) (model.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (mongo userRepository) Delete(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
