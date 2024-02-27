// @title GameStats API
// @description This is the GameStats API server.
// @version 1.0
// @BasePath /api

// @Security
// @SecurityDefinitions.basic BasicAuth
// -- @SecurityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	"GameStats.Api/internal/handlers"
	"GameStats.Infrastructure/tools"
	"fmt"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetReportCaller(true)

	config, err := tools.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	var r *chi.Mux = chi.NewRouter()

	database, err := tools.NewDatabase(config)
	if err != nil {
		log.Fatal(err)
	}
	userRepo := tools.NewUserRepo(database)

	var handler = handlers.Handler{
		UserRepo: userRepo,
	}

	(&handler).Handle(r)

	fmt.Printf("Starting API on port %v...", config.Server.Port)

	port := fmt.Sprintf(":%d", config.Server.Port)
	err = http.ListenAndServe(port, r)

	if err != nil {
		log.Fatal(err)
	}
}
