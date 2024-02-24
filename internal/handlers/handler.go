package handlers

import (
	"GameStats/internal/middleware"
	"GameStats/internal/tools"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

type Handler struct {
	UserRepo *tools.UserRepo
}

func (handler Handler) Handle(r *chi.Mux) {
	middlewareHandler := handler.CreateMiddlewareHandler()

	r.Use(chimiddleware.StripSlashes)

	r.Route("/test", func(r chi.Router) {
		r.Use(middlewareHandler.Authorization)

		r.Get("/super", handler.GetUserDetails)
	})
}

func (handler *Handler) CreateMiddlewareHandler() *middleware.Handler {
	return &middleware.Handler{
		UserRepo: handler.UserRepo,
	}
}
