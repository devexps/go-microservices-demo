package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-microservices-demo/app/user/service/internal/data"
	"github.com/devexps/go-microservices-demo/gen/api/go/common/pagination"
	v1 "github.com/devexps/go-microservices-demo/gen/api/go/user/service/v1"
)

// UserService .
type UserService struct {
	v1.UserServiceServer

	log *log.Helper
	uc  *data.UserRepo
}

// NewUserService .
func NewUserService(logger log.Logger, uc *data.UserRepo) *UserService {
	l := log.NewHelper(log.With(logger, "module", "user/service/user-service"))
	return &UserService{
		log: l,
		uc:  uc,
	}
}

// ListUser .
func (s *UserService) ListUser(ctx context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	return s.uc.ListUser(ctx, req)
}

// GetUser .
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	return s.uc.GetUser(ctx, req)
}

// CreateUser .
func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*emptypb.Empty, error) {
	return s.uc.CreateUser(ctx, req)
}

// UpdateUser .
func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*emptypb.Empty, error) {
	return s.uc.UpdateUser(ctx, req)
}

// DeleteUser .
func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	return s.uc.DeleteUser(ctx, req)
}

// VerifyPassword .
func (s *UserService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (*v1.User, error) {
	return s.uc.VerifyPassword(ctx, req)
}

// UserExists .
func (s *UserService) UserExists(ctx context.Context, req *v1.UserExistsRequest) (*v1.UserExistsResponse, error) {
	return s.uc.UserExists(ctx, req)
}
