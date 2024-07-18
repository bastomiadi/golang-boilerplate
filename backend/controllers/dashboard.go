package controllers

import (
	"html/template"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	// Parse main layout template and content template
	tmpl := template.Must(template.ParseFiles(
		"backend/views/partials/header.html",
		"backend/views/partials/sidebar.html",
		"backend/views/partials/footer.html",
		"backend/views/layouts/main.html",
		"backend/views/dashboard/index.html",
		"backend/views/categories/modals.html",
	))

	// Data for main layout template
	data := struct {
		Title string
	}{
		Title: "Dashboard",
	}

	// Execute main layout template with data
	if err := tmpl.ExecuteTemplate(w, "main.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
