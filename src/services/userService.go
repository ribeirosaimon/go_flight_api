package services

import (
	"github.com/pkg/errors"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/repository"
	"time"
)

const (
	_FIELD_ERROR          = "field can be blank"
	_FIELD_ERROR_SAVE     = "erro to save in repository"
	_PASSWORD_FIELD_ERROR = "password need more 6 char"
)

func SaveUserService(dto model.AccountDto) (model.Account, error) {
	var newAcc = model.Account{}
	if dto.Name == "" {
		return newAcc, errors.New(_FIELD_ERROR)
	}
	if len(dto.Password) < 6 {
		return newAcc, errors.New(_PASSWORD_FIELD_ERROR)
	}

	newAcc.Name = dto.Name
	newAcc.LastName = dto.LastName
	newAcc.Password = dto.Password
	newAcc.CreatedAt = time.Now()
	newAcc.UpdatedAt = time.Now()

	savedAccount, err := repository.Save(newAcc)
	if err != nil {
		return newAcc, errors.New(_FIELD_ERROR_SAVE)
	}

	return savedAccount.SanitizerAccount(), nil
}

func FindAllUserService() ([]model.Account, error) {
	all, err := repository.FindAll()
	if err != nil {
		return []model.Account{}, errors.New(err.Error())
	}
	return all, nil
}

func FindOneUserService(id string) (model.Account, error) {
	byId, err := repository.FindById(id)
	if err != nil {
		return model.Account{}, errors.New(err.Error())
	}
	return byId, nil
}

func UpdateUserService(id string, dto model.AccountDto) (model.Account, error) {
	update, err := repository.Update(id, dto)
	if err != nil {
		return model.Account{}, errors.New(err.Error())
	}
	return update, err
}

func DeleteUserService(id string) error {
	err := repository.Delete(id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
