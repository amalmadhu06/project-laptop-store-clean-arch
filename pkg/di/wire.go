//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/api"
	handler "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/api/handler"
	config "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/config"
	db "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/db"
	repository "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository"
	usecase "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
		handler.NewOtpHandler,
		usecase.NewOtpUseCase,
		repository.NewOtpRepository,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
