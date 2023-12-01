package server

import (
	"github.com/devexps/go-bootstrap"
	conf "github.com/devexps/go-bootstrap/gen/api/go/conf/v1"

	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware/logging"
	"github.com/devexps/go-micro/v2/transport/grpc"

	"github.com/devexps/go-microservices-demo/app/user/service/internal/service"
	v1 "github.com/devexps/go-microservices-demo/gen/api/go/user/service/v1"
)

// NewGRPCServer new an GRPC server.
func NewGRPCServer(
	cfg *conf.Bootstrap, logger log.Logger,
	userSvc *service.UserService,
) *grpc.Server {
	srv := bootstrap.CreateGrpcServer(cfg, logging.Server(logger))

	v1.RegisterUserServiceServer(srv, userSvc)

	return srv
}
