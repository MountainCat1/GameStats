package handlers

import (
	_ "GameStats.Api/docs"
	"GameStats.Api/internal/middleware"
	"GameStats.Domain/repositories"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	UserRepo repositories.IUserDetailsRepo
}

func (h Handler) Handle(r *chi.Mux) {
	middlewareHandler := h.CreateMiddlewareHandler()

	r.Use(chimiddleware.StripSlashes)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition
	))

	r.Route("/api", func(r chi.Router) {

		r.Route("/test", func(r chi.Router) {
			r.Use(middlewareHandler.Authorization)

			r.Get("/super", h.GetUserDetails)
		})
	})
}

func (h *Handler) CreateMiddlewareHandler() *middleware.Handler {
	return &middleware.Handler{
		UserRepo: h.UserRepo,
	}
}
