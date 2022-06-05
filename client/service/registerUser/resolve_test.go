package registerUser

import (
	"reflect"
	"testing"
	"transactions/model"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/require"
)

//go:generate moq -out mocks_test.go . Env

func TestResolve(t *testing.T) {
	userID := 1
	userName := "Dmitrii"

	type args struct {
		e     Env
		input model.RegisterUserInput
	}
	tests := []struct {
		name    string
		args    args
		want    model.RegisterUserOrErrorPayload
		wantErr bool
	}{
		{
			args: args{
				e: &EnvMock{
					CreateUserFunc: func(data *model.User) error {
						data.ID = userID
						data.Name = &userName
						return nil
					},
					CreateUserQueueFunc: func(userID int) (amqp.Queue, error) {
						return amqp.Queue{Name: "queue"}, nil
					},
					CreateConsumerForQueueFunc: func(qname string) {
						require.Equal(t, "queue", qname)
					},
				},
				input: model.RegisterUserInput{
					Name: &userName,
				},
			},
			want: model.RegisterUserPayload{User: &model.User{
				ID:   userID,
				Name: &userName,
			}},
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
