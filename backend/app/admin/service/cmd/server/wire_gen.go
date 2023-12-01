// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/devexps/go-bootstrap/gen/api/go/conf/v1"
	"github.com/devexps/go-micro/v2"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/registry"
	"github.com/devexps/go-microservices-demo/app/admin/service/internal/data"
	"github.com/devexps/go-microservices-demo/app/admin/service/internal/server"
	"github.com/devexps/go-microservices-demo/app/admin/service/internal/service"
)

// Injectors from wire.go:

// initApp init GoMicro application.
func initApp(logger log.Logger, registrar registry.Registrar, bootstrap *v1.Bootstrap) (*micro.App, func(), error) {
	authenticator := data.NewAuthenticator(bootstrap)
	authorizer := data.NewAuthorizer()
	discovery := data.NewDiscovery(bootstrap)
	userServiceClient := data.NewUserServiceClient(discovery, bootstrap)
	userTokenServiceClient := data.NewUserTokenServiceClient(discovery, bootstrap)
	authenticationService := service.NewAuthenticationService(logger, userServiceClient, userTokenServiceClient)
	userService := service.NewUserService(logger, userServiceClient)
	httpServer := server.NewHTTPServer(bootstrap, logger, authenticator, authorizer, authenticationService, userService)
	app := newApp(logger, registrar, httpServer)
	return app, func() {
	}, nil
}
