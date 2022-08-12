package repository

import (
	"context"
	"github.com/ribeirosaimon/go_flight_api/src/database"
	"github.com/ribeirosaimon/go_flight_api/src/models"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.GetCollection("Flight")
var ctx = context.Background()

func CreateFlight() {

}

func ReadFlights() (models.Flights, error) {
	var flights models.Flights
	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return flights, err
	}
	for cur.Next(ctx) {
		var flight models.Flight
		err = cur.Decode(&flight)
		if err != nil {
			return nil, err
		}

		flights = append(flights, &flight)
	}
	return flights, nil
}

func UpdateFlight() {

}

func DeleteFligh() {

}
