package graph

import (
	"transactions/model"
	"transactions/service/registerUser"
	"transactions/service/withdrawMoney"
)

//go:generate go run github.com/99designs/gqlgen generate
type Env interface {
	GetUserByID(id int) (*model.User, error)
	withdrawMoney.Env
	registerUser.Env
}

type Resolver struct {
	Env Env
}
