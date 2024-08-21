package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gprestore/gprestore-core/injector"
	"github.com/gprestore/gprestore-core/internal/config"
	"github.com/spf13/viper"
)

func main() {
	config.Load()

	route := injector.InjectRoute()
	route.Init()

	port := viper.GetString("server.port")
	log.Printf("Running on http://localhost:%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), route.Mux)
}
