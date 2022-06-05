package processTransaction

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
	"time"
	"transactions/model"

	"github.com/stretchr/testify/require"

	"github.com/streadway/amqp"
)

//go:generate moq -out mocks_test.go . Env

func TestResolve(t *testing.T) {
	userID := 1

	const layout = "Jan 2, 2006 at 3:04pm (PST)"

	timeNow, err := time.Parse(layout, "Feb 4, 2014 at 6:05pm (PST)")
	if err != nil {
		t.Fatal(err)
	}

	bodyWithReplenish, err := json.Marshal(model.Transaction{
		UserID:    userID,
		CreatedAt: timeNow,
		Operation: model.ReplenishPaymentType,
		Status:    model.ReceivedOperationStatus,
	})
	if err != nil {
		log.Fatal(err)
	}

	bodyWithWithdraw, err := json.Marshal(model.Transaction{
		UserID:    userID,
		CreatedAt: timeNow,
		Operation: model.WithdrawPaymentType,
		Status:    model.ReceivedOperationStatus,
	})
	if err != nil {
		log.Fatal(err)
	}

	type args struct {
		e       Env
		message amqp.Delivery
	}
	tests := []struct {
		name    string
		args    args
		want    model.Transaction
		wantErr bool
	}{
		{
			name: "nil user",
			args: args{
				e: &EnvMock{
					GetUserByIDFunc: func(id int) (*model.User, error) {
						return nil, nil
					},
				},
				message: amqp.Delivery{Body: bodyWithReplenish},
			},
			wantErr: true,
			want: model.Transaction{
				UserID:    userID,
				CreatedAt: timeNow,
				Status:    model.ReceivedOperationStatus,
				Operation: model.ReplenishPaymentType,
			},
		},
		{
			name: "success with ReplenishPaymentType",
			args: args{
				e: &EnvMock{
					GetUserByIDFunc: func(id int) (*model.User, error) {
						return &model.User{ID: userID}, nil
					},
					UpdateTransactionFunc: func(data *model.Transaction) error {
						require.Equal(t, userID, data.UserID)
						return nil
					},
					UpdateUserFunc: func(data *model.User) error {
						return nil
					},
					NowFunc: func() time.Time {
						return timeNow
					},
				},
				message: amqp.Delivery{Body: bodyWithReplenish},
			},
			wantErr: false,
			want: model.Transaction{
				UserID:     userID,
				CreatedAt:  timeNow,
				FinishedAt: timeNow,
				Status:     model.ProcessedOperationStatus,
				Operation:  model.ReplenishPaymentType,
			},
		},
		{
			name: "success with WithdrawPaymentType",
			args: args{
				e: &EnvMock{
					GetUserByIDFunc: func(id int) (*model.User, error) {
						return &model.User{ID: userID}, nil
					},
					UpdateTransactionFunc: func(data *model.Transaction) error {
						require.Equal(t, userID, data.UserID)
						return nil
					},
					UpdateUserFunc: func(data *model.User) error {
						return nil
					},
					NowFunc: func() time.Time {
						return timeNow
					},
				},
				message: amqp.Delivery{Body: bodyWithWithdraw},
			},
			wantErr: false,
			want: model.Transaction{
				UserID:     userID,
				CreatedAt:  timeNow,
				FinishedAt: timeNow,
				Status:     model.ProcessedOperationStatus,
				Operation:  model.WithdrawPaymentType,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Resolve(tt.args.e, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resolve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
