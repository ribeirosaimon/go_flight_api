package models

import "time"

type User struct {
	ID        uint64    `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name,omitempty"  bson:"name,omitempty"`
	Username  string    `json:"username,omitempty"  bson:"username,omitempty"`
	Email     string    `json:"email,omitempty"  bson:"email,omitempty"`
	Password  string    `json:"password,omitempty"  bson:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"  bson:"createdAt,omitempty"`
}
