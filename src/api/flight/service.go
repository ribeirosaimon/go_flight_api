package flight

import (
	"context"
	"errors"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/repository"
	"time"
)

type flightService struct {
	repository repository.FlightRepositoryImpl
}

func CreateFlightService() flightService {
	return flightService{repository: repository.FlightRepository()}
}

func (s flightService) findById(id string) (model.Flight, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	foudFlight, err := s.repository.FindById(ctx, id)

	if err != nil {
		return model.Flight{}, err
	}

	return foudFlight, nil
}

func (s flightService) saveFlight(flight model.Flight) (model.Flight, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	if flight.Price == 0 {
		return model.Flight{}, errors.New("price can't be null")
	}
	if len(flight.Origin) == 0 {
		return model.Flight{}, errors.New("you need a origin")
	}
	if len(flight.Destination) == 0 {
		return model.Flight{}, errors.New("you need a destination")
	}
	flight.CreatedAt = time.Now()
	return s.repository.Save(ctx, flight)
}

func (s flightService) findAllFlights() ([]model.Flight, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	return s.repository.FindAll(ctx)
}

func (s flightService) cheapFlight() (model.Flight, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	return s.repository.FindMoreCheapFlight(ctx)
}

func (s flightService) deleteById(id string) error {
	if id == "" {
		return errors.New("need id by delete")
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	return s.repository.DeleteById(ctx, id)
}

func (s flightService) getLastFlight() (model.Flight, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	return s.repository.GetLastFlight(ctx)
}

func (s flightService) searchFlight(search model.SearchFilter) (model.SearchFilterResult, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	return s.repository.SearchFlight(ctx, search)
}
