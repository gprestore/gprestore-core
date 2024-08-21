package route

import (
	"net/http"

	"github.com/gprestore/gprestore-core/internal/delivery/rest"
	"github.com/gprestore/gprestore-core/internal/delivery/rest/middleware"
)

type Route struct {
	Mux         *http.ServeMux
	Middleware  *middleware.Middleware
	UserHandler *rest.UserHandler
	AuthHandler *rest.AuthHandler
}

func New(
	mux *http.ServeMux,
	middleware *middleware.Middleware,
	userHandler *rest.UserHandler,
	authHandler *rest.AuthHandler,
) *Route {
	return &Route{
		Mux:         mux,
		Middleware:  middleware,
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}
}

func (r *Route) Init() {
	r.AuthRoutes()
	r.UserRoutes()
}

func (r *Route) AuthRoutes() {
	r.Mux.Handle("GET /auth/oauth/{provider}", r.Middleware.Guest(http.HandlerFunc(r.AuthHandler.OAuth)))
	r.Mux.Handle("GET /callback/oauth/{provider}", r.Middleware.Guest(http.HandlerFunc(r.AuthHandler.OAuthCallback)))
}

func (r *Route) UserRoutes() {
	r.Mux.Handle("POST /user", r.Middleware.Admin(http.HandlerFunc(r.UserHandler.CreateUser)))
	r.Mux.Handle("PATCH /user/{id}", r.Middleware.User(http.HandlerFunc(r.UserHandler.UpdateUserById)))
	r.Mux.Handle("DELETE /user/{id}", r.Middleware.User(http.HandlerFunc(r.UserHandler.DeleteUserById)))
	r.Mux.Handle("GET /users", r.Middleware.Admin(http.HandlerFunc(r.UserHandler.FindMany)))
	r.Mux.Handle("GET /user", r.Middleware.User(http.HandlerFunc(r.UserHandler.FindOne)))
}
