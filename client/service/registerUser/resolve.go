package registerUser

import (
	"transactions/errors"
	"transactions/model"

	"github.com/streadway/amqp"
)

type Env interface {
	CreateUser(data *model.User) error
	CreateUserQueue(userID int) (amqp.Queue, error)
	CreateConsumerForQueue(qname string)
}

//Resolve - создаёт пользователя очередь для него и слушателя очереди
func Resolve(e Env, input model.RegisterUserInput) (model.RegisterUserOrErrorPayload, error) {
	user := model.User{
		Name: input.Name,
	}

	err := e.CreateUser(&user)
	if err != nil {
		return errors.InternalErrorf("failed to CreateUser: %v", err)
	}

	q, err := e.CreateUserQueue(user.ID)
	if err != nil {
		return errors.InternalErrorf("failed to CreateUserQueue: %v", err)
	}

	go e.CreateConsumerForQueue(q.Name)

	return model.RegisterUserPayload{User: &user}, err
}
