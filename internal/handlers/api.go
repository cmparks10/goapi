package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	//Global Middleware

	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}