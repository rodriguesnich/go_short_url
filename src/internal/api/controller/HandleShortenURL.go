package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"url_short/internal/api/response"
	"url_short/internal/store"
)

type shortenURLRequest struct {
	URL string `json:"url"`
}

type shortenURLResponse struct {
	Code string `json:"code"`
}

func HandleShortenURL(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body shortenURLRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			response.SendJSON(w, response.Data{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if _, err := url.Parse(body.URL); err != nil {
			response.SendJSON(w, response.Data{Error: "invalid url passed"}, http.StatusBadRequest)
			return
		}

		code, err := store.SaveShortenedURL(r.Context(), body.URL)

		if err != nil {
			slog.Error("failed to create Code", "error", err)
			response.SendJSON(w, response.Data{Error: "something went wrong"}, http.StatusInternalServerError)
			return
		}

		response.SendJSON(w, response.Data{Data: shortenURLResponse{Code: code}}, http.StatusCreated)
	}
}
