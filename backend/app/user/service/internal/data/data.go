package data

import (
	"context"

	conf "github.com/devexps/go-bootstrap/gen/api/go/conf/v1"

	"github.com/devexps/go-microservices-demo/app/user/service/internal/data/ent"
	"github.com/devexps/go-microservices-demo/app/user/service/internal/data/ent/migrate"
	"github.com/devexps/go-utils/entgo"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/devexps/go-micro/v2/log"
)

// Data .
type Data struct {
	log *log.Helper
	db  *entgo.EntClient[*ent.Client]
}

// NewData .
func NewData(
	logger log.Logger,
	db *entgo.EntClient[*ent.Client],
) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/user-service"))

	d := &Data{
		log: l,
		db:  db,
	}
	return d, func() {
		l.Info("message", "closing the data resources")
		d.db.Close()
	}, nil
}

// NewEntClient .
func NewEntClient(cfg *conf.Bootstrap, logger log.Logger) *entgo.EntClient[*ent.Client] {
	l := log.NewHelper(log.With(logger, "module", "ent/data/user-service"))

	drv, err := entgo.CreateDriver(cfg.Data.Database.Driver, cfg.Data.Database.Source,
		int(cfg.Data.Database.MaxIdleConnections),
		int(cfg.Data.Database.MaxOpenConnections),
		cfg.Data.Database.ConnectionMaxLifetime.AsDuration(),
	)
	if err != nil {
		l.Fatalf("failed opening connection to db: %v", err)
		return nil
	}
	client := ent.NewClient(
		ent.Driver(drv),
		ent.Log(func(a ...any) {
			l.Debug(a...)
		}),
	)
	if cfg.Data.Database.Migrate {
		if err = client.Schema.Create(context.Background(), migrate.WithForeignKeys(true)); err != nil {
			l.Fatalf("failed creating schema resources: %v", err)
		}
	}
	return entgo.NewEntClient(client, drv)
}
