package repository

import (
	"context"
	"fmt"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/repository/users"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepo struct {
	writer *mongo.Collection
}

func NewMongoRepository(writer *mongo.Database, collection string) users.UserRepository {
	return &MongoRepo{writer: writer.Collection(collection)}
}

func (m MongoRepo) Create(ctx context.Context, newUser model.Account) error {
	one, err := m.writer.InsertOne(ctx, newUser)
	if err != nil {
		return err
	}
	fmt.Println(one)
	return nil
}

func (m MongoRepo) GetUserByEmail(ctx context.Context, email string) ([]model.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (m MongoRepo) GetByID(ctx context.Context, ID int64) (*model.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (m MongoRepo) GetAll(ctx context.Context) ([]model.Account, error) {
	//TODO implement me
	panic("implement me")
}
