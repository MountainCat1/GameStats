package handlers

import (
	"GameStats.Api/internal/middleware"
	"GameStats.Domain/repositories"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

type Handler struct {
	UserRepo repositories.IUserDetailsRepo
}

func (h Handler) Handle(r *chi.Mux) {
	middlewareHandler := h.CreateMiddlewareHandler()

	r.Use(chimiddleware.StripSlashes)

	r.Route("/test", func(r chi.Router) {
		r.Use(middlewareHandler.Authorization)

		r.Get("/super", h.GetUserDetails)
	})
}

func (h *Handler) CreateMiddlewareHandler() *middleware.Handler {
	return &middleware.Handler{
		UserRepo: h.UserRepo,
	}
}
