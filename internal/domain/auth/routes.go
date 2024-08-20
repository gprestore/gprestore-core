package auth

import (
	"net/http"

	"github.com/gprestore/gprestore-core/internal/middleware"
)

type Routes struct {
	mux     *http.ServeMux
	handler *Handler
}

func NewRoutes(mux *http.ServeMux, handler *Handler) *Routes {
	return &Routes{
		mux:     mux,
		handler: handler,
	}
}

func (r *Routes) Init() {
	r.mux.Handle("GET /callback/oauth", middleware.Guest(http.HandlerFunc(r.handler.Callback)))
	r.mux.Handle("GET /auth/login/google", middleware.Guest(http.HandlerFunc(r.handler.LoginGoogle)))
}
