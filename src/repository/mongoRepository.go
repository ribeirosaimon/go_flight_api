package repository

type container struct {
	Account AccountRepositoryImpl
	Flight  FlightRepositoryImpl
}

func NewMongoRepository() *container {
	return &container{
		Account: UserRepository(),
		Flight:  FlightRepository(),
	}
}
