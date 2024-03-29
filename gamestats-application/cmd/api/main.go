// @title Your API Title
// @description Your API Description
// @version 1.0
// @termsOfService http://your-terms-of-service-url.com

// @BasePath /api

// @SecurityDefinitions.basic basicAuth
// @in header
// @name Authentication

package main

import (
	"fmt"
	"gamestats-application/internal/handlers"
	"gamestats-intrastructure/infrastructure"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	fmt.Println(time.Now())

	log.SetReportCaller(true)

	config, err := infrastructure.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	var r *chi.Mux = chi.NewRouter()

	var handler = handlers.ConstructHandler(config)

	handler.Handle(r)

	log.Info("Starting API on port ", config.Server.Port, "...")

	port := fmt.Sprintf(":%d", config.Server.Port)
	err = http.ListenAndServe(port, r)

	if err != nil {
		log.Fatal(err)
	}
}
