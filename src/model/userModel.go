package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	USER  = "USER"
	ADMIN = "ADMIN"
)

type LoggedUser struct {
	Username string   `json:"username"`
	UserId   string   `json:"userId"`
	Roles    []string `json:"roles"`
}

type UserAccessToken struct {
	Token string `json:"access_token"`
}

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccountDto struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Account struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Username  string             `json:"username" bson:"username,omitempty"`
	Password  string             `json:"password" bson:"password"`
	Roles     []string           `json:"roles" bson:"roles"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

func (a *Account) SanitizerAccount() Account {
	return Account{a.ID,
		a.Name,
		a.LastName,
		a.Username,
		"",
		a.Roles,
		a.CreatedAt,
		a.UpdatedAt,
	}
}
