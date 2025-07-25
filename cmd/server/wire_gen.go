// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"grpc_gateway_framework/internal/app"
	"grpc_gateway_framework/internal/conf"
	"grpc_gateway_framework/internal/logger"
)

// Injectors from wire.go:

func initApp(appConfig *conf.AppConfig) (*app.App, func(), error) {
	int2 := appConfig.Port
	logConfig := appConfig.LogConfig
	string2 := appConfig.Mode
	zapLogger, err := logger.NewLogger(logConfig, string2)
	if err != nil {
		return nil, nil, err
	}
	v := conf.NewUnaryInterceptors()
	appApp, err := app.NewApp(int2, zapLogger, v)
	if err != nil {
		return nil, nil, err
	}
	return appApp, func() {
	}, nil
}
