package models

type Flight struct {
	ID          uint64            `bson:"id,omitempty"`
	Origin      map[string]string `bson:"origin,omitempty"`
	Destination map[string]string `bson:"destination,omitempty"`
	Month       string            `bson:"month,omitempty"`
	Travel      string            `bson:"travel,omitempty"`
	TimeGoing   uint64            `bson:"timeGoing,omitempty"`
	TimeReturn  uint64            `bson:"timeReturn,omitempty"`
	Price       uint64            `bson:"price,omitempty"`
	Airlines    []string          `bson:"airlines,omitempty"`
}
