package handlers

import (
	"GameStats/internal/tools"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

type Handler struct {
	UserRepo *tools.UserRepo
}

func (handler Handler) Handle(r *chi.Mux) {
	r.Use(chimiddleware.StripSlashes)

	r.Route("/test", func(r chi.Router) {
		//r.Use(middleware.Authorization)

		r.Get("/super", handler.GetUserDetails)
	})
}
