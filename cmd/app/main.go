package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gprestore/gprestore-core/internal/config"
	"github.com/gprestore/gprestore-core/internal/database"
	"github.com/gprestore/gprestore-core/internal/delivery/rest"
	"github.com/gprestore/gprestore-core/internal/delivery/rest/middleware"
	"github.com/gprestore/gprestore-core/internal/delivery/rest/route"
	"github.com/gprestore/gprestore-core/internal/repository"
	"github.com/gprestore/gprestore-core/internal/service"
	"github.com/spf13/viper"
)

func main() {
	config.Load()
	db := database.NewMongoDB()
	validate := validator.New()
	mux := http.NewServeMux()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, validate)
	userHandler := rest.NewUserHandler(userService)

	authService := service.NewAuthService(userRepository)
	authHandler := rest.NewAuthHandler(authService)

	middleware := middleware.NewMiddleware(authService)

	route := route.Route{
		Mux:         mux,
		Middleware:  middleware,
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}
	route.Init()

	port := viper.GetString("server.port")
	log.Printf("Running on http://localhost:%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
}
