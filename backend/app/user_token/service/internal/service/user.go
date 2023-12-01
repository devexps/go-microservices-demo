package service

import (
	"context"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-microservices-demo/app/user_token/service/internal/data"
	userV1 "github.com/devexps/go-microservices-demo/gen/api/go/user/service/v1"
	v1 "github.com/devexps/go-microservices-demo/gen/api/go/user_token/service/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UserTokenService .
type UserTokenService struct {
	v1.UserTokenServiceServer

	log *log.Helper
	uc  *data.UserTokenRepo
}

// NewUserTokenService .
func NewUserTokenService(logger log.Logger, uc *data.UserTokenRepo) *UserTokenService {
	l := log.NewHelper(log.With(logger, "module", "user/service/user_token-service"))
	return &UserTokenService{
		log: l,
		uc:  uc,
	}
}

// GenerateToken .
func (s *UserTokenService) GenerateToken(ctx context.Context, req *userV1.User) (*v1.TokenResponse, error) {
	return s.uc.GenerateToken(ctx, req)
}

// RemoveToken .
func (s *UserTokenService) RemoveToken(ctx context.Context, req *userV1.User) (*emptypb.Empty, error) {
	return s.uc.RemoveToken(ctx, req)
}

// GetRefreshToken .
func (s *UserTokenService) GetRefreshToken(ctx context.Context, req *userV1.User) (*v1.TokenResponse, error) {
	return s.uc.GetRefreshToken(ctx, req)
}

// GenerateAccessToken .
func (s *UserTokenService) GenerateAccessToken(ctx context.Context, req *userV1.User) (*v1.TokenResponse, error) {
	return s.uc.GenerateAccessToken(ctx, req)
}
