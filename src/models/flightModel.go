package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Flight struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Origin      map[string]string  `bson:"origin,omitempty"`
	Destination map[string]string  `bson:"destination,omitempty"`
	Month       string             `bson:"month,omitempty"`
	Travel      string             `bson:"travel,omitempty"`
	TimeGoing   uint64             `bson:"timeGoing,omitempty"`
	TimeReturn  uint64             `bson:"timeReturn,omitempty"`
	Price       uint64             `bson:"price,omitempty"`
	Airlines    [][]string         `bson:"airlines,omitempty"`
}

type Flights []*Flight
