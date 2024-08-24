//go:build wireinject
// +build wireinject

package injector

import (
	"net/http"

	"github.com/google/wire"
	"github.com/gprestore/gprestore-core/internal/delivery/mq"
	"github.com/gprestore/gprestore-core/internal/delivery/rest"
	"github.com/gprestore/gprestore-core/internal/delivery/rest/middleware"
	"github.com/gprestore/gprestore-core/internal/delivery/rest/route"
	"github.com/gprestore/gprestore-core/internal/infrastructure/database"
	"github.com/gprestore/gprestore-core/internal/repository"
	"github.com/gprestore/gprestore-core/internal/service"
	"github.com/gprestore/gprestore-core/internal/validation"
)

func InjectRoute() *route.Route {
	wire.Build(
		database.NewMongoDB,
		validation.New,
		http.NewServeMux,
		middleware.NewMiddleware,

		repository.NewUserRepository,
		service.NewUserService,
		rest.NewUserHandler,

		service.NewAuthService,
		rest.NewAuthHandler,

		repository.NewStoreRepository,
		service.NewStoreService,
		rest.NewStoreHandler,

		repository.NewItemRepository,
		service.NewItemService,
		rest.NewItemHandler,

		repository.NewStockRepository,
		service.NewStockService,
		rest.NewStockHandler,

		repository.NewOrderRepository,
		service.NewOrderService,
		rest.NewOrderHandler,

		service.NewPaymentService,
		rest.NewPaymentHandler,

		route.New,
	)

	return nil
}

func InjectConsumer() *mq.Consumer {
	wire.Build(
		service.NewMailService,
		mq.NewConsumer,
	)

	return nil
}
