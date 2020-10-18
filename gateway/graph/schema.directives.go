package graph

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jdc-lab/daily-shit/gateway/graph/generated"
	userService "github.com/jdc-lab/daily-shit/user-service/proto/user"
)

func Directives() generated.DirectiveRoot {
	return generated.DirectiveRoot{
		IsAdmin: func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
			claims, ok := ctx.Value("claims").(*userService.TokenClaims)

			if !ok || claims == nil || !claims.IsAdmin {
				// block calling the next resolver
				return nil, fmt.Errorf("Access denied")
			}

			// or let it pass through
			return next(ctx)
		},
		IsAuthenticated: func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
			claims, ok := ctx.Value("claims").(*userService.TokenClaims)

			if !ok || claims == nil {
				// block calling the next resolver
				return nil, fmt.Errorf("Access denied")
			}

			// or let it pass through
			return next(ctx)
		},
	}
}
