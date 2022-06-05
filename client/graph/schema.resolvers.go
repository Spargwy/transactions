package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"transactions/graph/generated"
	"transactions/model"
	"transactions/service/registerUser"
	"transactions/service/replenishTheBalance"
	"transactions/service/withdrawMoney"
)

func (r *mutationResolver) WithdrawMoney(ctx context.Context, input model.WithdrawMoneyInput) (model.WithdrawMoneyOrErrorPayload, error) {
	return withdrawMoney.Resolve(r.Env, input)
}

func (r *mutationResolver) ReplenishTheBalance(ctx context.Context, input model.ReplenishTheBalanceInput) (model.ReplenishTheBalanceOrErrorPayload, error) {
	return replenishTheBalance.Resolve(r.Env, input)
}

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUserInput) (model.RegisterUserOrErrorPayload, error) {
	return registerUser.Resolve(r.Env, input)
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	return r.Env.GetUserByID(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
