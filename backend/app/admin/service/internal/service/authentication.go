package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/devexps/go-micro/v2/log"

	v1 "github.com/devexps/go-microservices-demo/gen/api/go/admin/service/v1"
	userV1 "github.com/devexps/go-microservices-demo/gen/api/go/user/service/v1"
	userTokenV1 "github.com/devexps/go-microservices-demo/gen/api/go/user_token/service/v1"
	"github.com/devexps/go-microservices-demo/pkg/middleware/auth"
)

// AuthenticationService .
type AuthenticationService struct {
	v1.AuthenticationServiceHTTPServer

	log *log.Helper
	uc  userV1.UserServiceClient
	utc userTokenV1.UserTokenServiceClient
}

// NewAuthenticationService .
func NewAuthenticationService(logger log.Logger, uc userV1.UserServiceClient, utc userTokenV1.UserTokenServiceClient) *AuthenticationService {
	l := log.NewHelper(log.With(logger, "module", "authentication/service/admin-service"))
	return &AuthenticationService{
		log: l,
		uc:  uc,
		utc: utc,
	}
}

// Login .
func (s *AuthenticationService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	var user *userV1.User
	var err error
	if user, err = s.uc.VerifyPassword(ctx, &userV1.VerifyPasswordRequest{
		UserName: req.GetUsername(),
		Password: req.GetPassword(),
	}); err != nil {
		return nil, err
	}
	tokenRsp, err := s.utc.GenerateToken(ctx, user)
	if err != nil {
		return nil, err
	}
	return &v1.LoginResponse{
		GrandType:    "bearer",
		AccessToken:  tokenRsp.GetAccessToken(),
		RefreshToken: tokenRsp.GetRefreshToken(),
	}, nil
}

// Logout .
func (s *AuthenticationService) Logout(ctx context.Context, req *v1.LogoutRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	req.Id = authInfo.UserId

	_, err = s.utc.RemoveToken(ctx, &userV1.User{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// RefreshToken .
func (s *AuthenticationService) RefreshToken(ctx context.Context, req *v1.RefreshTokenRequest) (*v1.RefreshTokenResponse, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	rsp, err := s.utc.GetRefreshToken(ctx, &userV1.User{
		Id: authInfo.UserId,
	})
	if err != nil {
		return nil, err
	}
	if rsp.GetRefreshToken() != req.GetRefreshToken() {
		return nil, v1.ErrorInvalidToken("invalid refresh token")
	}
	rsp2, err := s.utc.GenerateAccessToken(ctx, &userV1.User{
		Id: authInfo.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RefreshTokenResponse{
		GrandType:    "bearer",
		RefreshToken: rsp.GetRefreshToken(),
		AccessToken:  rsp2.GetAccessToken(),
	}, nil
}

// GetMe .
func (s *AuthenticationService) GetMe(ctx context.Context, req *v1.GetMeRequest) (*userV1.User, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	req.Id = authInfo.UserId

	userRsp, err := s.uc.GetUser(ctx, &userV1.GetUserRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return userRsp, nil
}
