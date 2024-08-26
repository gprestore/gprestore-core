package route

import (
	"net/http"

	"github.com/gprestore/gprestore-core/internal/delivery/rest"
	"github.com/gprestore/gprestore-core/internal/delivery/rest/middleware"
)

type Route struct {
	Mux            *http.ServeMux
	Middleware     *middleware.Middleware
	UserHandler    *rest.UserHandler
	AuthHandler    *rest.AuthHandler
	StoreHandler   *rest.StoreHandler
	ItemHandler    *rest.ItemHandler
	StockHandler   *rest.StockHandler
	OrderHandler   *rest.OrderHandler
	PaymentHandler *rest.PaymentHandler
}

func New(
	mux *http.ServeMux,
	middleware *middleware.Middleware,
	userHandler *rest.UserHandler,
	authHandler *rest.AuthHandler,
	storeHandler *rest.StoreHandler,
	itemHandler *rest.ItemHandler,
	stockHandler *rest.StockHandler,
	orderHandler *rest.OrderHandler,
	paymentHandler *rest.PaymentHandler,
) *Route {
	return &Route{
		Mux:            mux,
		Middleware:     middleware,
		UserHandler:    userHandler,
		AuthHandler:    authHandler,
		StoreHandler:   storeHandler,
		ItemHandler:    itemHandler,
		StockHandler:   stockHandler,
		OrderHandler:   orderHandler,
		PaymentHandler: paymentHandler,
	}
}

func (r *Route) Init() {
	r.AuthRoutes()
	r.UserRoutes()
	r.StoreRoutes()
	r.ItemRoutes()
	r.StockRoutes()
	r.OrderRoutes()
	r.PaymentRoutes()
}

func (r *Route) AuthRoutes() {
	r.Mux.Handle("GET /auth/oauth/{provider}", r.Middleware.Guest(http.HandlerFunc(r.AuthHandler.OAuth)))
	r.Mux.Handle("GET /callback/oauth/{provider}", r.Middleware.Guest(http.HandlerFunc(r.AuthHandler.OAuthCallback)))
	r.Mux.Handle("GET /auth/token/refresh", r.Middleware.Guest(http.HandlerFunc(r.AuthHandler.CheckRefreshToken)))
}

func (r *Route) UserRoutes() {
	r.Mux.Handle("POST /user", r.Middleware.Admin(http.HandlerFunc(r.UserHandler.Create)))
	r.Mux.Handle("PATCH /user/{id}", r.Middleware.User(http.HandlerFunc(r.UserHandler.UpdateById)))
	r.Mux.Handle("DELETE /user/{id}", r.Middleware.User(http.HandlerFunc(r.UserHandler.DeleteById)))
	r.Mux.Handle("GET /users", r.Middleware.Admin(http.HandlerFunc(r.UserHandler.FindMany)))
	r.Mux.Handle("GET /user", r.Middleware.User(http.HandlerFunc(r.UserHandler.FindOne)))
}

func (r *Route) StoreRoutes() {
	r.Mux.Handle("POST /store", r.Middleware.User(http.HandlerFunc(r.StoreHandler.Create)))
	r.Mux.Handle("PATCH /store/{id}", r.Middleware.User(http.HandlerFunc(r.StoreHandler.UpdateById)))
	r.Mux.Handle("DELETE /store/{id}", r.Middleware.User(http.HandlerFunc(r.StoreHandler.DeleteById)))
	r.Mux.Handle("GET /stores", r.Middleware.Guest(http.HandlerFunc(r.StoreHandler.FindMany)))
	r.Mux.Handle("GET /store", r.Middleware.Guest(http.HandlerFunc(r.StoreHandler.FindOne)))
}

func (r *Route) ItemRoutes() {
	r.Mux.Handle("POST /item", r.Middleware.User(http.HandlerFunc(r.ItemHandler.Create)))
	r.Mux.Handle("PATCH /item/{id}", r.Middleware.User(http.HandlerFunc(r.ItemHandler.UpdateById)))
	r.Mux.Handle("DELETE /item/{id}", r.Middleware.User(http.HandlerFunc(r.ItemHandler.DeleteById)))
	r.Mux.Handle("GET /items", r.Middleware.Guest(http.HandlerFunc(r.ItemHandler.FindMany)))
	r.Mux.Handle("GET /item", r.Middleware.Guest(http.HandlerFunc(r.ItemHandler.FindOne)))
}

func (r *Route) StockRoutes() {
	r.Mux.Handle("POST /stock", r.Middleware.User(http.HandlerFunc(r.StockHandler.Create)))
	r.Mux.Handle("PATCH /stock/{id}", r.Middleware.User(http.HandlerFunc(r.StockHandler.UpdateById)))
	r.Mux.Handle("GET /stocks", r.Middleware.Admin(http.HandlerFunc(r.StockHandler.FindOne)))
}

func (r *Route) OrderRoutes() {
	r.Mux.Handle("POST /order", r.Middleware.Guest(http.HandlerFunc(r.OrderHandler.Create)))
	r.Mux.Handle("PATCH /order/{id}", r.Middleware.Admin(http.HandlerFunc(r.OrderHandler.UpdateById)))
	r.Mux.Handle("DELETE /order/{id}", r.Middleware.Admin(http.HandlerFunc(r.OrderHandler.DeleteById)))
	r.Mux.Handle("GET /orders", r.Middleware.User(http.HandlerFunc(r.OrderHandler.FindMany)))
	r.Mux.Handle("GET /order", r.Middleware.Guest(http.HandlerFunc(r.OrderHandler.FindOne)))
}

func (r *Route) PaymentRoutes() {
	r.Mux.Handle("GET /payment/channels", r.Middleware.Guest(http.HandlerFunc(r.PaymentHandler.FindPaymentChannels)))
}
