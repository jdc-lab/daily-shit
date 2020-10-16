package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtAuthenticator struct {
	secret []byte
}

type JwtClaims struct {
	IsAdmin bool   `json:"isAdmin"`
	UserId  string `json:"userId"`
	Expires int64  `json:"expires"`
}

func (c JwtClaims) Valid() error {
	now := jwt.TimeFunc().Unix()

	if now > c.Expires {
		return jwt.ValidationError{
			Inner:  errors.New("token is expired"),
			Errors: jwt.ValidationErrorExpired,
		}
	}

	return nil
}

func (a jwtAuthenticator) NewToken(ctx context.Context, userId string) (string, error) {
	claims := JwtClaims{
		IsAdmin: true,
		UserId:  userId,
		Expires: time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(a.secret)
	if err != nil {
		return "", fmt.Errorf("could not sign jwt token\n%w", err)
	}

	return tokenString, nil
}

func (a jwtAuthenticator) Validate(ctx context.Context, tokenString string) (JwtClaims, error) {
	var claims JwtClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		// Validate expected algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return a.secret, nil
	})

	if err != nil {
		return JwtClaims{}, fmt.Errorf("could not parse claims:\n%w", err)
	}

	fmt.Println(token.Claims)

	if err := claims.Valid(); err != nil {
		return JwtClaims{}, fmt.Errorf("invalid claims:\n%w", err)
	}

	return claims, nil
}
