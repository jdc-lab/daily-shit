package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jdc-lab/daily-shit/gateway/graph/generated"
	"github.com/jdc-lab/daily-shit/gateway/graph/model"
	userService "github.com/jdc-lab/daily-shit/user-service/proto/user"
	"github.com/opentracing/opentracing-go/log"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user model.NewUser) (*model.CreateUserResponse, error) {
	claims := ctx.Value("claims").(*userService.TokenClaims)

	u := userService.CreateUserRequest{
		Claims:   claims,
		IsAdmin:  user.IsAdmin,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	createReq, err := r.UserService.Create(ctx, &u)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %w", err)
	}

	return &model.CreateUserResponse{
		ID: createReq.Id,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, user model.LoginUser) (*model.LoginUserResponse, error) {
	a := userService.AuthRequest{
		Username: user.Username,
		Password: user.Password,
	}

	authReq, err := r.UserService.Auth(ctx, &a)
	if err != nil {
		log.Error(fmt.Errorf("could not authenticate user %w", err))
		// client should not get the information why exactly
		return nil, fmt.Errorf("could not authenticate user")
	}

	return &model.LoginUserResponse{
		ID:    authReq.Id,
		Token: authReq.Token,
	}, nil
}

func (r *queryResolver) User(_ context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
