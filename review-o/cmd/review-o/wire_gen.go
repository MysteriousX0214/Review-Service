// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"review-o/internal/biz"
	"review-o/internal/conf"
	"review-o/internal/data"
	"review-o/internal/server"
	"review-o/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	reviewClient := data.NewReviewServiceClient(discovery)
	dataData, cleanup, err := data.NewData(confData, reviewClient, logger)
	if err != nil {
		return nil, nil, err
	}
	operationRepo := data.NewOperationRepo(dataData, logger)
	operationUsecase := biz.NewOperationUsecase(operationRepo, logger)
	operationService := service.NewOperationService(operationUsecase)
	grpcServer := server.NewGRPCServer(confServer, operationService, logger)
	httpServer := server.NewHTTPServer(confServer, operationService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
