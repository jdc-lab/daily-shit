package main

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type jwtAuthenticator struct {
	secret []byte
}

func (a jwtAuthenticator) NewToken(ctx context.Context, userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(a.secret)
	if err != nil {
		return "", fmt.Errorf("could not sign jwt token %w", err)
	}

	return tokenString, nil
}

func (a jwtAuthenticator) Validate(ctx context.Context, token string) {
	panic("implement me")
}
