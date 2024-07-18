package controllers

import (
	"html/template"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	layout := "backend/views/layouts/main.html"
	dashboard := "backend/views/dashboard.html"
	tmpl, err := template.ParseFiles(layout, dashboard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "main", nil)
}
