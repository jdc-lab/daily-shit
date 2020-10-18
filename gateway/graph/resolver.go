package graph

import (
	userService "github.com/jdc-lab/daily-shit/user-service/proto/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService userService.UserServiceClient
}
