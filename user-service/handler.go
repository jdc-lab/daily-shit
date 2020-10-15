package main

import (
	"context"

	pb "./proto/user"
)

type repository interface {
	Create(ctx context.Context, username string, email string, password string) (string, error)
	Get(ctx context.Context, id string) (user, error)
}

type handler struct {
	repo repository
}

func (h handler) Create(ctx context.Context, u *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res := new(pb.CreateUserResponse)
	if id, err := h.repo.Create(ctx, u.Username, u.Email, u.Password); err != nil {
		res.Errors = append(res.Errors, &pb.Error{
			Code:        1,
			Description: err.Error(),
		})
	} else {
		res.Id = id
	}

	return res, nil
}

func (h handler) Get(ctx context.Context, u *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res := new(pb.GetUserResponse)
	if user, err := h.repo.Get(ctx, u.Id); err != nil {
		res.Errors = append(res.Errors, &pb.Error{
			Code:        1,
			Description: err.Error(),
		})
	} else {
		res.Id = user.id
		res.Email = user.email
		res.Username = user.username
	}

	return res, nil
}

func (h handler) Auth(ctx context.Context, request *pb.AuthRequest) (*pb.AuthResponse, error) {
	panic("implement me")
}

func (h handler) ValidateToken(ctx context.Context, request *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	panic("implement me")
}
