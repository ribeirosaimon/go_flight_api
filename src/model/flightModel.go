package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Flight struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Origin      map[string]string  `json:"origin" bson:"origin"`
	Destination map[string]string  `json:"destination" bson:"destination"`
	TravelAt    string             `json:"travelAt" bson:"travelAt"`
	TimeGoing   uint               `json:"timeGoing" bson:"timeGoing"`
	TimeReturn  uint               `json:"timeReturn" bson:"timeReturn"`
	Price       uint64             `json:"price" bson:"price"`
	Airlines    []string           `json:"airlines" bson:"airlines"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
}
