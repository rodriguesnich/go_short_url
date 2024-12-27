package api

import (
	"url_short/internal/api/controller"
	"url_short/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func AddRoutes(r *chi.Mux, store store.Store) {

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.SetHeader("contenContent-Type", "application/json"))

		r.Route("/url", func(r chi.Router) {
			r.Post("/shorten", controller.HandleShortenURL(store))
			r.Get("/{code}", controller.HandleGetShortenedURL(store))
		})
	})

	r.Get("/{code}", controller.HandleRedirectToURL(store))
}
