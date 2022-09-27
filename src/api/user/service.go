package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/repository"
	"github.com/ribeirosaimon/go_flight_api/src/security"
)

type userService struct {
	repository repository.AccountRepositoryImpl
}

func UserService() userService {
	return userService{repository: repository.UserRepository()}
}

func (s userService) FindAllUserService() ([]model.Account, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	all, err := s.repository.FindAll(ctx)
	if err != nil {
		return []model.Account{}, err
	}
	return all, nil
}

func (s userService) FindOneUserService(id string) (model.Account, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	byId, err := s.repository.FindById(ctx, id)
	if err != nil {
		return model.Account{}, errors.New(err.Error())
	}
	return byId, nil
}

func (s userService) UpdateUserService(ID string, dto model.AccountDto) (model.Account, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	newAccount, err := s.FindOneUserService(ID)
	if err != nil {
		return model.Account{}, err
	}

	if dto.Name != "" {
		newAccount.Name = dto.Name
	}
	if dto.Username != "" {
		newAccount.Name = dto.Username
	}
	if dto.LastName != "" {
		newAccount.Name = dto.LastName
	}

	if dto.Password != "" {
		encriptedPassword, err := security.EncriptyPassword(dto.Password)
		if err != nil {
			return model.Account{}, err
		}
		newAccount.Password = string(encriptedPassword)
	}

	update, err := s.repository.Update(ctx, newAccount)
	if err != nil {
		return model.Account{}, errors.New(err.Error())
	}
	return update, err
}

func (s userService) DeleteUserService(id string) error {

	err := s.repository.Delete(context.Background(), id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (s userService) promotedToAdmin(loggedUser model.LoggedUser, id string) error {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	userDb, err := s.repository.FindById(ctx, id)
	if err != nil {
		return err
	}

	userDb.Roles = append(userDb.Roles, model.ADMIN)

	_, err = repository.NewMongoRepository().Account.Update(ctx, userDb)

	if err != nil {
		return err
	}

	return nil
}

func (s userService) verifyPassword(loggedUser model.LoggedUser, user model.Account) (bool, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	foundUser, err := s.repository.FindById(ctx, loggedUser.UserId)

	if err != nil {
		return false, err
	}

	if err := security.VerifyPassword(foundUser.Password, user.Password); err != nil {
		return false, err
	}
	return true, err
}
