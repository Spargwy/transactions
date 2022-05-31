package graph

import "transactions/model"

//go:generate go run github.com/99designs/gqlgen generate
type Env interface {
	GetUserByID(id int) (*model.User, error)
}

type Resolver struct {
	Env Env
}
