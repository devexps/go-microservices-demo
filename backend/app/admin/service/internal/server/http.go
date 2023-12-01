package server

import (
	"context"

	"github.com/devexps/go-bootstrap"
	conf "github.com/devexps/go-bootstrap/gen/api/go/conf/v1"
	swaggerUI "github.com/devexps/go-swagger-ui"

	"github.com/devexps/go-micro/v2/middleware"
	"github.com/devexps/go-micro/v2/middleware/authn"
	"github.com/devexps/go-micro/v2/middleware/authz"
	"github.com/devexps/go-micro/v2/middleware/logging"
	"github.com/devexps/go-micro/v2/middleware/selector"

	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/transport/http"

	authnEngine "github.com/devexps/go-micro/v2/middleware/authn/engine"
	authzEngine "github.com/devexps/go-micro/v2/middleware/authz/engine"

	"github.com/devexps/go-microservices-demo/app/admin/service/cmd/server/assets"
	"github.com/devexps/go-microservices-demo/app/admin/service/internal/service"
	v1 "github.com/devexps/go-microservices-demo/gen/api/go/admin/service/v1"
	"github.com/devexps/go-microservices-demo/pkg/middleware/auth"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	cfg *conf.Bootstrap, logger log.Logger,
	authenticator authnEngine.Authenticator, authorizer authzEngine.Authorizer,
	authnSvc *service.AuthenticationService,
	userSvc *service.UserService,
) *http.Server {
	srv := bootstrap.CreateHTTPServer(cfg, newHTTPMiddleware(authenticator, authorizer, logger)...)

	v1.RegisterAuthenticationServiceHTTPServer(srv, authnSvc)
	v1.RegisterUserServiceHTTPServer(srv, userSvc)

	if cfg.GetServer().GetHttp().GetEnableSwagger() {
		swaggerUI.RegisterSwaggerUIServerWithOption(
			srv,
			swaggerUI.WithTitle("GoMicro Monolithic demo"),
			swaggerUI.WithMemoryData(assets.OpenApiData, "yaml"),
		)
	}
	return srv
}

// newHTTPMiddleware .
func newHTTPMiddleware(authenticator authnEngine.Authenticator, authorizer authzEngine.Authorizer, logger log.Logger) []middleware.Middleware {
	var ms []middleware.Middleware
	ms = append(ms, logging.Server(logger))
	ms = append(ms, selector.Server(
		authn.Server(authenticator),
		auth.Server(),
		authz.Server(authorizer),
	).Match(newHTTPWhiteListMatcher()).Build())
	return ms
}

// newHTTPWhiteListMatcher .
func newHTTPWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	whiteList[v1.OperationAuthenticationServiceLogin] = true
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
