//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	"github.com/devexps/go-micro/v2"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/registry"

	conf "github.com/devexps/go-bootstrap/gen/api/go/conf/v1"

	"github.com/devexps/go-microservices-demo/app/user_token/service/internal/data"
	"github.com/devexps/go-microservices-demo/app/user_token/service/internal/server"
	"github.com/devexps/go-microservices-demo/app/user_token/service/internal/service"
)

// initApp init GoMicro application.
func initApp(log.Logger, registry.Registrar, *conf.Bootstrap) (*micro.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, data.ProviderSet, newApp))
}
