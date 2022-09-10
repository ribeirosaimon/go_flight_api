package repository

type container struct {
	Account mongoRepository
}

func NewMongoRepository() *container {
	return &container{Account: newUserRepository()}
}
