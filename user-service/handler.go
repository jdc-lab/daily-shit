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
	Validate(ctx context.Context, tokenString string) (JwtClaims, error)
}

type handler struct {
	pb.UnimplementedUserServiceServer
	repo repository
	auth authenticator
}

func (h handler) Create(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res := new(pb.CreateUserResponse)
	if id, err := h.repo.Create(ctx, request.GetUsername(), request.GetEmail(), request.GetPassword()); err != nil {
		res.Errors = append(res.Errors, &pb.Error{
			Code:        1,
			Description: err.Error(),
		})
	} else {
		res.Id = id
	}

	return res, nil
}

func (h handler) Get(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res := new(pb.GetUserResponse)
	if user, err := h.repo.Get(ctx, request.GetId()); err != nil {
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
	user, err := h.repo.GetByName(ctx, request.GetUsername())
	if err != nil {
		return nil, fmt.Errorf("could not get name\n%w", err)
	}

	bcrypt.CompareHashAndPassword([]byte(user.passwordHash), []byte(request.GetPassword()))
	if err != nil {
		return nil, fmt.Errorf("could not validate password\n%w", err)
	}

	token, err := h.auth.NewToken(ctx, user.id)
	if err != nil {
		return nil, fmt.Errorf("could not generate token\n%w", err)
	}

	return &pb.AuthResponse{
		Id:    user.id,
		Token: token,
	}, nil
}

func (h handler) ValidateToken(ctx context.Context, request *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	claims, err := h.auth.Validate(ctx, request.GetToken())

	if err != nil {
		return nil, fmt.Errorf("invalid token\n%w", err)
	}

	return &pb.ValidateTokenResponse{
		IsAdmin: claims.IsAdmin,
		UserId:  claims.UserId,
		Expires: claims.Expires,
	}, nil
}
