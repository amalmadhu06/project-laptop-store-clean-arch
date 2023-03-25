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
		//database connection
		db.ConnectDatabase,

		//handler
		handler.NewAdminHandler,
		handler.NewUserHandler,
		handler.NewOtpHandler,
		handler.NewProductHandler,
		handler.NewCartHandler,
		handler.NewOrderHandler,
		handler.NewPaymentHandler,
		handler.NewWishlistHandler,

		//database queries
		repository.NewAdminRepository,
		repository.NewUserRepository,
		repository.NewOtpRepository,
		repository.NewProductRepository,
		repository.NewCartRepository,
		repository.NewOrderRepository,
		repository.NewPaymentRepository,
		repository.NewWishlistRepository,

		//use case
		usecase.NewAdminUseCase,
		usecase.NewUserUseCase,
		usecase.NewOtpUseCase,
		usecase.NewProductUseCase,
		usecase.NewCartUseCase,
		usecase.NewOrderUseCase,
		usecase.NewPaymentUseCase,
		usecase.NewWishlistUsecase,

		//server connection
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
