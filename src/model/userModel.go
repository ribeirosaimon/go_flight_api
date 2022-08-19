package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AccountDto struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Password string `json:"password"`
}

type Account struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

func (a *Account) SanitizerAccount() Account {
	return Account{a.ID,
		a.Name,
		a.LastName,
		"",
		a.CreatedAt,
		a.UpdatedAt,
	}
}