package route

import (
	"net/http"

	"github.com/gprestore/gprestore-core/internal/delivery/rest"
	"github.com/gprestore/gprestore-core/internal/delivery/rest/middleware"
)

type Route struct {
	Mux          *http.ServeMux
	Middleware   *middleware.Middleware
	UserHandler  *rest.UserHandler
	AuthHandler  *rest.AuthHandler
	StoreHandler *rest.StoreHandler
	ItemHandler  *rest.ItemHandler
}

func New(
	mux *http.ServeMux,
	middleware *middleware.Middleware,
	userHandler *rest.UserHandler,
	authHandler *rest.AuthHandler,
	storeHandler *rest.StoreHandler,
	itemHandler *rest.ItemHandler,
) *Route {
	return &Route{
		Mux:          mux,
		Middleware:   middleware,
		UserHandler:  userHandler,
		AuthHandler:  authHandler,
		StoreHandler: storeHandler,
		ItemHandler:  itemHandler,
	}
}

func (r *Route) Init() {
	r.AuthRoutes()
	r.UserRoutes()
	r.StoreRoutes()
	r.ItemRoutes()
}

func (r *Route) AuthRoutes() {
	r.Mux.Handle("GET /auth/oauth/{provider}", r.Middleware.Guest(http.HandlerFunc(r.AuthHandler.OAuth)))
	r.Mux.Handle("GET /callback/oauth/{provider}", r.Middleware.Guest(http.HandlerFunc(r.AuthHandler.OAuthCallback)))
	r.Mux.Handle("GET /auth/token/refresh", r.Middleware.Guest(http.HandlerFunc(r.AuthHandler.CheckRefreshToken)))
}

func (r *Route) UserRoutes() {
	r.Mux.Handle("POST /user", r.Middleware.Admin(http.HandlerFunc(r.UserHandler.CreateUser)))
	r.Mux.Handle("PATCH /user/{id}", r.Middleware.User(http.HandlerFunc(r.UserHandler.UpdateUserById)))
	r.Mux.Handle("DELETE /user/{id}", r.Middleware.User(http.HandlerFunc(r.UserHandler.DeleteUserById)))
	r.Mux.Handle("GET /users", r.Middleware.Admin(http.HandlerFunc(r.UserHandler.FindMany)))
	r.Mux.Handle("GET /user", r.Middleware.User(http.HandlerFunc(r.UserHandler.FindOne)))
}

func (r *Route) StoreRoutes() {
	r.Mux.Handle("POST /store", r.Middleware.User(http.HandlerFunc(r.StoreHandler.CreateStore)))
	r.Mux.Handle("PATCH /store/{id}", r.Middleware.User(http.HandlerFunc(r.StoreHandler.UpdateStoreById)))
	r.Mux.Handle("DELETE /store/{id}", r.Middleware.User(http.HandlerFunc(r.StoreHandler.DeleteStoreById)))
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
