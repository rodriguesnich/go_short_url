package app

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func GetMainPage() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tmplPath := filepath.Join("public", "views", "main.html")

		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "Failed to parse template", http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(w, "index", struct {
			HasError bool
		}{false})
		if err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
			return
		}
	}
}
