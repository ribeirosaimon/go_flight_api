package services

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/repository"
	"github.com/ribeirosaimon/go_flight_api/src/security"
	"time"
)

const contextTime = 2

type userService struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func UserService() userService {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*contextTime)
	return userService{ctx: ctx, cancel: cancelFunc}
}

func (u userService) SaveOneAccount(account model.AccountDto) (model.Account, error) {
	defer u.cancel()
	var newAcc model.Account

	if account.Name == "" {
		return newAcc, errors.New("You need send Name")
	}
	if account.Password == "" {
		return newAcc, errors.New("You need send Password")
	}

	password, err := security.EncriptyPassword(account.Password)
	if err != nil {
		return newAcc, errors.New("Error to encrypt password")
	}

	newAcc.Name = account.Name
	newAcc.LastName = account.LastName
	newAcc.Username = account.Username
	newAcc.Password = string(password)
	newAcc.CreatedAt = time.Now()
	newAcc.UpdatedAt = time.Now()
	newAcc.Roles = []string{model.USER}
	mongoRepository := repository.NewMongoRepository()

	return mongoRepository.Account.Save(u.ctx, newAcc)

}

func (u userService) FindAllUserService() ([]model.Account, error) {
	all, err := repository.NewMongoRepository().Account.FindAll(u.ctx)
	if err != nil {
		return []model.Account{}, err
	}
	return all, nil
}

func (u userService) FindOneUserService(id string) (model.Account, error) {
	byId, err := repository.NewMongoRepository().Account.FindById(u.ctx, id)
	if err != nil {
		return model.Account{}, errors.New(err.Error())
	}
	return byId, nil
}

func (u userService) UpdateUserService(id string, dto model.AccountDto) (model.Account, error) {
	update, err := repository.NewMongoRepository().Account.Update(u.ctx, id, dto)
	if err != nil {
		return model.Account{}, errors.New(err.Error())
	}
	return update, err
}

func (u userService) DeleteUserService(id string) error {
	err := repository.NewMongoRepository().Account.Delete(u.ctx, id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (u userService) UserLogin(dto model.LoginDto) (model.UserAccessToken, error) {
	defer u.cancel()
	var accessToken model.UserAccessToken

	if dto.Username == "" {
		return model.UserAccessToken{}, errors.New("Username can't be null")
	}
	account, err := repository.NewMongoRepository().Account.FindByUsername(u.ctx, dto.Username)
	if err != nil {
		return model.UserAccessToken{}, errors.New("user not found")
	}

	if err := security.VerifyPassword(account.Password, dto.Password); err != nil {
		return accessToken, errors.New("password as incorrect")
	}
	token, err := security.CreateToken(account)
	if err != nil {
		return accessToken, err
	}
	accessToken.Token = token
	return accessToken, err
}
