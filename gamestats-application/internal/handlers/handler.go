package handlers

import (
	_ "gamestats-application/docs"
	"gamestats-application/internal/middleware"
	domain "gamestats-domain"
	"gamestats-intrastructure/infrastructure"
	"gamestats-intrastructure/repositories"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	UserRepo  domain.UserRepo
	MatchRepo domain.MatchRepo
}

func ConstructHandler(config infrastructure.Config) *Handler {

	database, err := infrastructure.NewNoSqlDatabase(config)
	if err != nil {
		log.Fatal(err)
	}
	userRepo := infrastructure.NewUserRepository(database)

	matchRepo := repositories.MewMatchRepository(database)

	return &Handler{
		UserRepo:  userRepo,
		MatchRepo: matchRepo,
	}
}

func (h *Handler) Handle(r *chi.Mux) {
	middlewareHandler := h.CreateMiddlewareHandler()

	r.Use(chimiddleware.StripSlashes)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition
	))

	r.Route("/api", func(r chi.Router) {

		r.Route("/users", func(r chi.Router) {
			r.Post("/", h.CreateUser)
			r.With(middlewareHandler.Authentication).Get("/", h.GetUserDetails)
		})

		r.Route("/matches", func(r chi.Router) {
			r.Post("/", h.CreateMatch)
			r.With(middlewareHandler.Authentication).Post("/", h.CreateMatch)

			r.Get("/", h.GetMatches)
			r.With(middlewareHandler.Authentication).Get("/", h.GetMatches)
		})
	})
}

func (h *Handler) CreateMiddlewareHandler() *middleware.Handler {
	return &middleware.Handler{
		UserRepo: h.UserRepo,
	}
}
