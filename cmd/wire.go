//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"microdule-layout/pkg"
)

func InitService() (*pkg.App, error) {
	wire.Build(
		pkg.ProviderGlobal,
		pkg.ProviderRedis,
		pkg.ProviderDb,
		pkg.ProviderETCD,
		pkg.ProviderWeb,

		NewService,
	)
	return &pkg.App{}, nil
}
