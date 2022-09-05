package repository

import (
	"context"
	"github.com/ribeirosaimon/go_flight_api/src/model"
)

type mongoRepository interface {
	Save(ctx context.Context, account model.Account) (model.Account, error)
	FindById(ctx context.Context, ID string) (model.Account, error)
	FindByUsername(ctx context.Context, username string) (model.Account, error)
	FindAll(ctx context.Context) ([]model.Account, error)
	Update(ctx context.Context, ID string, account model.AccountDto) (model.Account, error)
	Delete(ctx context.Context, ID string) error
}
