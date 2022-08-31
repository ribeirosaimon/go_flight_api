package users

import (
	"context"
	"github.com/ribeirosaimon/go_flight_api/src/model"
)

type UserRepository interface {
	Create(ctx context.Context, newUser model.Account) error
	GetUserByEmail(ctx context.Context, email string) ([]model.Account, error)
	GetByID(ctx context.Context, ID int64) (*model.Account, error)
	GetAll(ctx context.Context) ([]model.Account, error)
}
