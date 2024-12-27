package controller

import (
	"errors"
	"log/slog"
	"net/http"
	"url_short/internal/api/response"
	"url_short/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

type getShortenedURLResponse struct {
	FullURL string `json:"full_url"`
}

func HandleGetShortenedURL(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		fullURL, err := store.GetFullURL(r.Context(), code)

		if err != nil {
			if errors.Is(err, redis.Nil) {
				response.SendJSON(w, response.Data{Error: "code not found"}, http.StatusNotFound)
				return
			}

			slog.Error("failed to retrieve Code", "error", err)
			response.SendJSON(w, response.Data{Error: "something went wrong"}, http.StatusInternalServerError)
			return
		}

		response.SendJSON(w, response.Data{Data: getShortenedURLResponse{FullURL: fullURL}}, http.StatusOK)
	}
}
