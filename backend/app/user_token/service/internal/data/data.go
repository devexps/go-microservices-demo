package data

import (
	"github.com/devexps/go-bootstrap"
	conf "github.com/devexps/go-bootstrap/gen/api/go/conf/v1"
	"github.com/devexps/go-micro/middleware/authn/engine/jwt/v2"
	authnEngine "github.com/devexps/go-micro/v2/middleware/authn/engine"
	"github.com/redis/go-redis/v9"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/devexps/go-micro/v2/log"
)

// Data .
type Data struct {
	log           *log.Helper
	rdb           *redis.Client
	authenticator authnEngine.Authenticator
}

// NewData .
func NewData(
	logger log.Logger,
	rdb *redis.Client,
	authenticator authnEngine.Authenticator,
) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/user-token-service"))

	d := &Data{
		log:           l,
		rdb:           rdb,
		authenticator: authenticator,
	}
	return d, func() {
		l.Info("message", "closing the data resources")
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewRedisClient creates redis client
func NewRedisClient(cfg *conf.Bootstrap, _ log.Logger) *redis.Client {
	return bootstrap.NewRedisClient(cfg.Data)
}

// NewAuthenticator .
func NewAuthenticator(cfg *conf.Bootstrap) authnEngine.Authenticator {
	authenticator, _ := jwt.NewAuthenticator(
		jwt.WithKey([]byte(cfg.Server.Grpc.Middleware.Auth.Key)),
		jwt.WithSigningMethod(cfg.Server.Grpc.Middleware.Auth.Method),
	)
	return authenticator
}
