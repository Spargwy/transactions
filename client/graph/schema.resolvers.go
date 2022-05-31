package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	generated1 "transactions/graph/generated"
	"transactions/model"
)

func (r *mutationResolver) WithdrawMoney(ctx context.Context, input *model.WithdrawMoneyInput) (model.WithdrawMoneyOrErrorPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ReplenishTheBalance(ctx context.Context, input *model.ReplenishTheBalanceInput) (model.ReplenishTheBalanceOrErrorPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	return r.Env.GetUserByID(id)
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
