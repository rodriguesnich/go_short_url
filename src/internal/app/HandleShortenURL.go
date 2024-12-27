package app

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"path/filepath"
	"url_short/internal/store"
)

func HandleShortenURL(store store.Store) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tmplPath := filepath.Join("public", "views", "main.html")

		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "Failed to parse template", http.StatusInternalServerError)
			return
		}

		url := r.FormValue("url")
		if url == "" {
			err = tmpl.ExecuteTemplate(w, "createForm", struct {
				HasError bool
			}{true})

			if err != nil {
				http.Error(w, "Failed to render template", http.StatusInternalServerError)
				return
			}

			return
		}

		code, err := store.SaveShortenedURL(r.Context(), url)

		if err != nil {
			slog.Error("failed to create Code", "error", err)
			return
		}

		baseURL := fmt.Sprintf("%s://%s", r.URL.Scheme, r.Host)

		err = tmpl.ExecuteTemplate(w, "shortUrlResult", struct {
			Code    string
			BaseURL string
		}{code, baseURL})

		if err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
			return
		}
	}
}
