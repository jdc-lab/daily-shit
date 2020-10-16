package main

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	id           string
	username     string
	email        string
	passwordHash string
}

type inMemoryRepository struct {
	users       map[string]*user
	usersByName map[string]*user
}

func (r *inMemoryRepository) init() {
	if r.users == nil {
		r.users = make(map[string]*user)
	}

	if r.usersByName == nil {
		r.usersByName = make(map[string]*user)
	}
}

func (r *inMemoryRepository) Create(_ context.Context, username string, email string, password string) (string, error) {
	r.init()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	u := user{
		id:           uuid.New().String(),
		username:     username,
		email:        email,
		passwordHash: string(hash),
	}

	r.users[u.id] = &u
	r.usersByName[u.username] = &u

	return u.id, nil
}

func (r *inMemoryRepository) Get(_ context.Context, id string) (user, error) {
	r.init()

	if u, ok := r.users[id]; ok {
		return *u, nil
	}

	return user{}, errors.New("user not found")
}

func (r *inMemoryRepository) GetByName(_ context.Context, username string) (user, error) {
	r.init()

	if u, ok := r.usersByName[username]; ok {
		return *u, nil
	}

	return user{}, errors.New("user not found")
}
