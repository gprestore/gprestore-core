package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gprestore/gprestore-core/internal/config"
	"github.com/gprestore/gprestore-core/internal/database"
	"github.com/gprestore/gprestore-core/internal/domain/auth"
	"github.com/gprestore/gprestore-core/internal/domain/user"
	"github.com/spf13/viper"
)

func main() {
	config.Load()
	db := database.NewMongoDB()
	validate := validator.New()
	mux := http.NewServeMux()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository, validate)
	userHandler := user.NewHandler(userService)
	user.NewRoutes(mux, userHandler).Init()

	authHandler := auth.NewHandler(userService)
	auth.NewRoutes(mux, authHandler).Init()

	port := viper.GetString("server.port")
	log.Printf("Running on http://localhost:%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
}
