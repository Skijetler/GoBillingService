// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Skijetler/GoBillingService/auth/internal/biz"
	"github.com/Skijetler/GoBillingService/auth/internal/conf"
	"github.com/Skijetler/GoBillingService/auth/internal/data"
	"github.com/Skijetler/GoBillingService/auth/internal/pkg/hash"
	"github.com/Skijetler/GoBillingService/auth/internal/pkg/paseto"
	"github.com/Skijetler/GoBillingService/auth/internal/server"
	"github.com/Skijetler/GoBillingService/auth/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, hasher *conf.Hasher, tokenMaker *conf.TokenMaker, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewAuthRepo(dataData, logger)
	passwordHasher := hash.NewPasswordHasher(hasher)
	pasetoTokenMaker, err := paseto.NewPasetoMaker(tokenMaker)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	authUsecase := biz.NewAuthUsecase(userRepo, logger, passwordHasher, pasetoTokenMaker)
	authService := service.NewAuthService(authUsecase)
	grpcServer := server.NewGRPCServer(confServer, authService, logger)
	httpServer := server.NewHTTPServer(confServer, authService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
