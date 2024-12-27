package app

import (
	"net/http"
	"url_short/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func AddRoutes(r *chi.Mux, store store.Store) {

	fs := http.FileServer(http.Dir("public"))
	r.Mount("/public/", http.StripPrefix("/public/", fs))

	r.Route("/app", func(r chi.Router) {
		r.Use(middleware.SetHeader("contenContent-Type", "text/html"))

		r.Get("/", GetMainPage())
		r.Post("/shorten", HandleShortenURL(store))
	})

}
