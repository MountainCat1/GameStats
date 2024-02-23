package main

import (
	"GameStats/internal/handlers"
	"fmt"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()

	handlers.Handle(r)

	fmt.Print("Starting API on port 8080...")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}
}
