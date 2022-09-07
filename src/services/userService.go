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

func SaveOneAccount(account model.AccountDto) (model.Account, error) {
	var newAcc model.Account

	if account.Name == "" {
		return newAcc, errors.New("You need send Name")
	}
	if account.Password == "" {
		return newAcc, errors.New("You need send Password")
	}

	newAcc.Name = account.Name
	newAcc.LastName = account.LastName
	newAcc.Username = account.Username
	newAcc.Password = account.Password
	newAcc.CreatedAt = time.Now()
	newAcc.UpdatedAt = time.Now()
	newAcc.Roles = []string{model.USER}

	return repository.NewMongoRepository().Account.Save(context.Background(), newAcc)

}

func FindAllUserService() ([]model.Account, error) {

	all, err := repository.NewMongoRepository().Account.FindAll(context.Background())
	if err != nil {
		return []model.Account{}, err
	}
	return all, nil
}

func FindOneUserService(id string) (model.Account, error) {

	byId, err := repository.NewMongoRepository().Account.FindById(context.Background(), id)
	if err != nil {
		return model.Account{}, errors.New(err.Error())
	}
	return byId, nil
}

func UpdateUserService(id string, dto model.AccountDto) (model.Account, error) {
	update, err := repository.NewMongoRepository().Account.Update(context.Background(), id, dto)
	if err != nil {
		return model.Account{}, errors.New(err.Error())
	}
	return update, err
}

func DeleteUserService(id string) error {

	err := repository.NewMongoRepository().Account.Delete(context.Background(), id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func UserLogin(dto model.LoginDto) (model.UserAccessToken, error) {
	var accessToken model.UserAccessToken

	if dto.Username == "" {
		return model.UserAccessToken{}, errors.New("Username can't be null")
	}
	account, err := repository.NewMongoRepository().Account.FindByUsername(context.Background(), dto.Username)
	if err != nil {
		return model.UserAccessToken{}, err
	}

	if err != nil {
		panic(err)
	}

	if err := security.VerifyPassword(account.Password, dto.Password); err != nil {
		return accessToken, err
	}
	token, err := security.CreateToken(account)
	if err != nil {
		return accessToken, err
	}
	accessToken.Token = token
	return accessToken, err
}
