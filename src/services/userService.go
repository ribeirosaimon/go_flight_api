package services

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/repository"
	"github.com/ribeirosaimon/go_flight_api/src/security"
	"time"
)

type userService struct {
	context context.Context
}

func (s userService) SaveOneAccount(account model.Account) (model.Account, error) {
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
	newAcc.Password = string(password)
	newAcc.CreatedAt = time.Now()
	newAcc.UpdatedAt = time.Now()
	newAcc.Roles = []string{model.USER}
	mongoRepository := repository.NewMongoRepository()
	return mongoRepository.Account.Save(s.context, newAcc)

}

//
//func FindAllUserService() ([]model.Account, error) {
//	all, err := repository.FindAll()
//	if err != nil {
//		return []model.Account{}, errors.New(err.Error())
//	}
//	return all, nil
//}
//
//func FindOneUserService(id string) (model.Account, error) {
//	byId, err := repository.FindById(id)
//	if err != nil {
//		return model.Account{}, errors.New(err.Error())
//	}
//	return byId, nil
//}
//
//func UpdateUserService(id string, dto model.AccountDto) (model.Account, error) {
//	update, err := repository.Update(id, dto)
//	if err != nil {
//		return model.Account{}, errors.New(err.Error())
//	}
//	return update, err
//}
//
//func DeleteUserService(id string) error {
//	err := repository.Delete(id)
//	if err != nil {
//		return errors.New(err.Error())
//	}
//	return nil
//}
//
//func UserLogin(dto model.LoginDto) (model.UserAccessToken, error) {
//	var accessToken model.UserAccessToken
//	if dto.Username == "" {
//		return model.UserAccessToken{}, errors.New("Username can't be null")
//	}
//	account, err := repository.FindUserByUsername("admin")
//	if err != nil {
//		return model.UserAccessToken{}, errors.New("user not found")
//	}
//
//	if err := security.VerifyPassword(account.Password, dto.Password); err != nil {
//		return accessToken, errors.New("password as incorrect")
//	}
//	token, err := security.CreateToken(account)
//	if err != nil {
//		return accessToken, err
//	}
//	accessToken.Token = token
//	return accessToken, err
//}
