package handlers

import (
	"GameStats/internal/middleware"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

func Handle(r *chi.Mux) {
	r.Use(chimiddleware.StripSlashes)

	r.Route("/test", func(r chi.Router) {
		r.Use(middleware.Authorization)

		r.Get("/super", GetUserDetails)
	})
}
