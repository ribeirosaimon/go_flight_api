package repository

import (
	"github.com/ribeirosaimon/go_flight_api/src/repository/users"
	"go.mongodb.org/mongo-driver/mongo"
)

type Container struct {
	User users.UserRepository
}
type Options struct {
	WriteMongo   *mongo.Client
	DatabaseName string
	Collection   string
}

func New(opts Options) *Container {
	return &Container{User: NewMongoRepository(opts.WriteMongo.Database(opts.DatabaseName), opts.Collection)}
}
