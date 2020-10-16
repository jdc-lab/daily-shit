package main

import (
	"context"
	"fmt"

	pb "./proto/user"
	"golang.org/x/crypto/bcrypt"
)

type repository interface {
	Create(ctx context.Context, username string, email string, password string) (string, error)
	Get(ctx context.Context, id string) (user, error)
	GetByName(ctx context.Context, username string) (user, error)
}

type authenticator interface {
	NewToken(ctx context.Context, userId string) (string, error)
	Validate(ctx context.Context, token string)
}

type handler struct {
	repo repository
	auth authenticator
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
	user, err := h.repo.GetByName(ctx, request.Username)
	if err != nil {
		return nil, fmt.Errorf("could not get name %w", err)
	}

	bcrypt.CompareHashAndPassword([]byte(user.passwordHash), []byte(request.Password))
	if err != nil {
		return nil, fmt.Errorf("could not validate password %w", err)
	}

	token, err := h.auth.NewToken(ctx, user.id)
	if err != nil {
		return nil, fmt.Errorf("could not generate token %w", err)
	}

	return &pb.AuthResponse{
		Id:    user.id,
		Token: token,
	}, nil
}

func (h handler) ValidateToken(ctx context.Context, request *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	panic("implement me")
}
