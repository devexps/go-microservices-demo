package data

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/devexps/go-bootstrap"

	conf "github.com/devexps/go-bootstrap/gen/api/go/conf/v1"

	userV1 "github.com/devexps/go-microservices-demo/gen/api/go/user/service/v1"
	userTokenV1 "github.com/devexps/go-microservices-demo/gen/api/go/user_token/service/v1"
	"github.com/devexps/go-microservices-demo/pkg/service"

	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/registry"

	"github.com/devexps/go-micro/middleware/authn/engine/jwt/v2"
	authnEngine "github.com/devexps/go-micro/v2/middleware/authn/engine"
	authzEngine "github.com/devexps/go-micro/v2/middleware/authz/engine"
	"github.com/devexps/go-micro/v2/middleware/authz/engine/noop"
)

// Data .
type Data struct {
	log             *log.Helper
	authenticator   authnEngine.Authenticator
	authorizer      authzEngine.Authorizer
	userClient      userV1.UserServiceClient
	userTokenClient userTokenV1.UserTokenServiceClient
}

// NewData .
func NewData(
	logger log.Logger,
	authenticator authnEngine.Authenticator,
	authorizer authzEngine.Authorizer,
	userClient userV1.UserServiceClient,
	userTokenClient userTokenV1.UserTokenServiceClient,
) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/admin-service"))

	d := &Data{
		log:             l,
		authenticator:   authenticator,
		authorizer:      authorizer,
		userClient:      userClient,
		userTokenClient: userTokenClient,
	}
	return d, func() {
		l.Info("message", "closing the data resources")
	}, nil
}

// NewAuthenticator .
func NewAuthenticator(cfg *conf.Bootstrap) authnEngine.Authenticator {
	authenticator, _ := jwt.NewAuthenticator(
		jwt.WithKey([]byte(cfg.Server.Http.Middleware.Auth.Key)),
		jwt.WithSigningMethod(cfg.Server.Http.Middleware.Auth.Method),
	)
	return authenticator
}

// NewAuthorizer .
func NewAuthorizer() authzEngine.Authorizer {
	return noop.NewAuthorizer()
}

// NewDiscovery .
func NewDiscovery(cfg *conf.Bootstrap) registry.Discovery {
	return bootstrap.NewConsulRegistry(cfg.Registry)
}

// NewUserServiceClient .
func NewUserServiceClient(r registry.Discovery, cfg *conf.Bootstrap) userV1.UserServiceClient {
	return userV1.NewUserServiceClient(bootstrap.CreateGrpcClient(context.Background(), r, service.UserService, cfg))
}

// NewUserTokenServiceClient .
func NewUserTokenServiceClient(r registry.Discovery, cfg *conf.Bootstrap) userTokenV1.UserTokenServiceClient {
	return userTokenV1.NewUserTokenServiceClient(bootstrap.CreateGrpcClient(context.Background(), r, service.UserTokenService, cfg))
}
