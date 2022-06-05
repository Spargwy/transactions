package withdrawMoney

import (
	"encoding/json"
	"strconv"
	"transactions/errors"
	"transactions/model"
)

type Env interface {
	GetUserByID(id int) (*model.User, error)
	CreateTransaction(data *model.Transaction) error
	PublishRMQMessage(qName string, body []byte) error
}

func Resolve(e Env, input model.WithdrawMoneyInput) (model.WithdrawMoneyOrErrorPayload, error) {
	if input.Amount <= 0 {
		return model.ErrorPayload{Message: "Amount can not be equal or smaller than 0"}, nil
	}

	user, err := e.GetUserByID(input.UserID)
	if err != nil {
		return errors.InternalErrorf("failed to GetUserByID: %v", err)
	}

	if user == nil {
		return model.ErrorPayload{Message: "User not found"}, nil
	}

	if user.Balance < input.Amount {
		return model.ErrorPayload{Message: "Not enough funds"}, nil
	}

	transaction := model.Transaction{
		UserID:        user.ID,
		Operation:     model.WithdrawPaymentType,
		BalanceBefore: user.Balance,
		Amount:        input.Amount,
	}

	err = e.CreateTransaction(&transaction)
	if err != nil {
		return errors.InternalErrorf("failed to create transaction: %v", err)
	}

	body, err := json.Marshal(transaction)
	if err != nil {
		return errors.InternalErrorf("failed to Marshal: %v", err)
	}

	err = e.PublishRMQMessage(strconv.Itoa(user.ID), body)

	return model.WithdrawMoneyPayload{
		AmountWrittenOff: input.Amount,
		UserID:           user.ID,
	}, err
}
