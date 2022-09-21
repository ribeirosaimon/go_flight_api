package login

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/repository"
	"github.com/ribeirosaimon/go_flight_api/src/security"
	"time"
)

type serviceLogin struct {
	repository repository.AccountRepositoryImpl
}

func ServiceLogin() serviceLogin {
	return serviceLogin{repository: repository.UserRepository()}
}

func (s serviceLogin) SaveOneAccount(account model.AccountDto) (model.Account, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
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

	return s.repository.Save(ctx, newAcc)
}

func (s serviceLogin) UserLogin(dto model.LoginDto) (model.UserAccessToken, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	var accessToken model.UserAccessToken

	if dto.Username == "" {
		return model.UserAccessToken{}, errors.New("Username can't be null")
	}

	account, err := s.repository.FindByUsername(ctx, dto.Username)

	if err != nil {
		return model.UserAccessToken{}, err
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

func ValidateLoggedUser(ctx *fiber.Ctx, searchId string) (model.LoggedUser, error) {
	val := ctx.Locals("loggedUser")
	loggedUser := val.(model.LoggedUser)

	if loggedUser.UserId != searchId {
		return model.LoggedUser{}, errors.New("you not have permission")
	}
	return loggedUser, nil
}

func WhoIsMe(ctx *fiber.Ctx) model.LoggedUser {
	val := ctx.Locals("loggedUser")
	return val.(model.LoggedUser)
}
