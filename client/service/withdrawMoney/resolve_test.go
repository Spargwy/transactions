package withdrawMoney

import (
	"reflect"
	"strconv"
	"testing"
	"transactions/model"

	"github.com/stretchr/testify/require"
)

//go:generate moq -out mocks_test.go . Env

func TestResolve(t *testing.T) {
	userID := 1

	type args struct {
		e     Env
		input model.WithdrawMoneyInput
	}
	tests := []struct {
		name    string
		args    args
		want    model.WithdrawMoneyOrErrorPayload
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				e: &EnvMock{
					GetUserByIDFunc: func(id int) (*model.User, error) {
						return &model.User{ID: userID, Balance: 100}, nil
					},
					CreateTransactionFunc: func(data *model.Transaction) error {
						require.Equal(t, data.UserID, userID)
						return nil
					},
					PublishRMQMessageFunc: func(qName string, body []byte) error {
						require.Equal(t, strconv.Itoa(userID), qName)
						return nil
					},
				},
				input: model.WithdrawMoneyInput{
					UserID: userID,
					Amount: 100,
				},
			},
			want: model.WithdrawMoneyPayload{
				UserID:           userID,
				AmountWrittenOff: 100,
			},
		},
		{
			name: "0 amount",
			args: args{
				input: model.WithdrawMoneyInput{
					UserID: userID,
					Amount: 0,
				},
			},
			want: model.ErrorPayload{Message: "Amount can not be equal or smaller than 0"},
		},
		{
			name: "nil user",
			args: args{
				e: &EnvMock{
					GetUserByIDFunc: func(id int) (*model.User, error) {
						return nil, nil
					},
				},
				input: model.WithdrawMoneyInput{
					UserID: userID,
					Amount: 100,
				},
			},
			want: model.ErrorPayload{Message: "User not found"},
		},
		{
			name: "not enough funds",
			args: args{
				e: &EnvMock{
					GetUserByIDFunc: func(id int) (*model.User, error) {
						return &model.User{ID: userID, Balance: 50}, nil
					},
					CreateTransactionFunc: func(data *model.Transaction) error {
						require.Equal(t, data.UserID, userID)
						return nil
					},
					PublishRMQMessageFunc: func(qName string, body []byte) error {
						require.Equal(t, strconv.Itoa(userID), qName)
						return nil
					},
				},
				input: model.WithdrawMoneyInput{
					UserID: userID,
					Amount: 100,
				},
			},
			want: model.ErrorPayload{
				Message: "Not enough funds",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Resolve(tt.args.e, tt.args.input)
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
