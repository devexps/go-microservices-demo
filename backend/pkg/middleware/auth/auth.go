package auth

import (
	"context"
	"strconv"

	"github.com/devexps/go-micro/v2/middleware"
	"github.com/devexps/go-micro/v2/middleware/authn"
	"github.com/devexps/go-micro/v2/middleware/authz"
	"github.com/devexps/go-micro/v2/transport"
)

var (
	action  = "ANY"
	project = "*"
)

type authInfo struct {
	UserId uint32
}

// Server .
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}
			authnClaims, ok := authn.FromContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}
			sub := authnClaims.GetSubject()
			path := tr.Operation()
			authzClaims := AuthzClaims{
				Subject:  sub,
				Action:   action,
				Resource: path,
				Project:  project,
			}
			ctx = authz.NewContext(ctx, &authzClaims)
			return handler(ctx, req)
		}
	}
}

// FromContext .
func FromContext(ctx context.Context) (*authInfo, error) {
	claims, ok := authn.FromContext(ctx)
	if !ok {
		return nil, ErrMissingJwtToken
	}
	userId, err := strconv.ParseUint(claims.GetSubject(), 10, 32)
	if err != nil {
		return nil, ErrExtractSubjectFailed
	}
	return &authInfo{
		UserId: uint32(userId),
	}, nil
}
