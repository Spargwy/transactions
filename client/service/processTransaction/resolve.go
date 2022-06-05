package processTransaction

import (
	"encoding/json"
	"fmt"
	"time"
	"transactions/model"

	"github.com/streadway/amqp"
)

type Env interface {
	UpdateTransaction(data *model.Transaction) error
	GetUserByID(id int) (*model.User, error)
	UpdateUser(data *model.User) error
	Now() time.Time
}

func Resolve(e Env, message amqp.Delivery) (model.Transaction, error) {
	var transaction model.Transaction
	err := json.Unmarshal(message.Body, &transaction)
	if err != nil {
		return transaction, fmt.Errorf("failed to Unmarshal body: %v", err)
	}

	user, err := e.GetUserByID(transaction.UserID)
	if err != nil {
		return transaction, fmt.Errorf("failed to GetUserByID: %v", err)
	}

	if user == nil {
		return transaction, fmt.Errorf("user not found")
	}

	switch transaction.Operation {
	case model.ReplenishPaymentType:
		user.Balance += transaction.Amount
	case model.WithdrawPaymentType:
		user.Balance -= transaction.Amount
	}

	err = e.UpdateUser(user)
	if err != nil {
		return transaction, fmt.Errorf("failed to UpdateUser: %v", err)
	}

	transaction.Status = model.ProcessedOperationStatus
	transaction.FinishedAt = e.Now()

	err = e.UpdateTransaction(&transaction)

	return transaction, err
}
