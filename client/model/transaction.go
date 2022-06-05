package model

import "time"

type OperationType string
type OperationStatus string

const (
	WithdrawPaymentType  OperationType = "withdraw"
	ReplenishPaymentType OperationType = "replenish"
)

const (
	ReceivedOperationStatus OperationStatus = "received"
	// In progress статус должен быть.
	// В нашем случае он пока не пригождается.
	// Но если добавляется обработка третьими лицами - must have
	InProgressOperationStatus OperationStatus = "in progress"
	ProcessedOperationStatus  OperationStatus = "processed"
	FailedOperationStatus     OperationStatus = "failed"
)

//Для сохранения истории транзакций
type Transaction struct {
	ID            int
	UserID        int
	Operation     OperationType
	BalanceBefore int
	BalanceAfter  int
	Amount        int
	Status        OperationStatus
	//Разделяем на время получения и время обработки,
	//чтобы, в случае чего, могли видеть, сколько времени
	//было потрачено
	CreatedAt  time.Time
	FinishedAt time.Time
}
