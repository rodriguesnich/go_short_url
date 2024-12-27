package controller

import (
	"net/http"
	"url_short/internal/store"

	"github.com/go-chi/chi/v5"
)

func HandleRedirectToURL(store store.Store) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		redirectURL, err := store.GetFullURL(r.Context(), code)

		if err != nil {
			http.Redirect(w, r, "", http.StatusNotFound)
		}

		http.Redirect(w, r, redirectURL, http.StatusFound)
	}
}
