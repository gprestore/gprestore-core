package user

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
	r.mux.Handle("POST /user", middleware.Admin(http.HandlerFunc(r.handler.CreateUser)))
	r.mux.Handle("PATCH /user/{id}", middleware.User(http.HandlerFunc(r.handler.UpdateUserById)))
	r.mux.Handle("DELETE /user/{id}", middleware.User(http.HandlerFunc(r.handler.DeleteUserById)))
	r.mux.Handle("GET /users", middleware.Admin(http.HandlerFunc(r.handler.FindMany)))
	r.mux.Handle("GET /user", middleware.Admin(http.HandlerFunc(r.handler.FindOne)))
}
