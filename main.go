package main

import (
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/repository"
	"time"
)

func main() {
	var acc = model.Account{
		Name:      "Nome",
		LastName:  "Ribeiro",
		Password:  "234",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	repository.Update("62fe3ce443dddd9fcdc10d11", acc)

}
