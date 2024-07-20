// backend/controllers/category.go
package controllers

import (
	"encoding/json"
	"golang-boilerplate/backend/models"
	"golang-boilerplate/config"
	"html/template"
	"log"
	"net/http"
)

func HandleCategoryIndex(w http.ResponseWriter, r *http.Request) {

	// Fetch menu items from the database
	menuItems, err := models.GetMenuItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse main layout template and content template
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"hasChildren": models.HasChildren,
		"dict":        models.Dict,
		"isActive":    models.IsActive,
	}).ParseFiles(
		"backend/views/partials/header.html",
		"backend/views/partials/sidebar.html",
		"backend/views/partials/footer.html",
		"backend/views/layouts/main.html",
		"backend/views/categories/index.html",
		"backend/views/categories/modals.html",
	))

	// Data for main layout template
	data := struct {
		Title      string
		MenuItems  []models.Menu
		CurrentURL string
	}{
		Title:      "Categories",
		MenuItems:  menuItems,
		CurrentURL: r.URL.Path,
	}

	// Execute main layout template with data
	if err := tmpl.ExecuteTemplate(w, "main.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandleCategoryList handles listing categories.
func HandleCategoryList(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	categories := []models.Category{}

	rows, err := db.Query("SELECT id, name FROM categories ORDER BY id DESC")
	if err != nil {
		log.Fatalf("Failed to query categories: %v", err)
		http.Error(w, "Failed to query categories", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			log.Fatalf("Failed to scan category: %v", err)
			http.Error(w, "Failed to scan categories", http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
		http.Error(w, "Error iterating over categories", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

// HandleCategoryCreate handles creating a new category.
func HandleCategoryCreate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	name := r.FormValue("name")

	_, err := db.Exec("INSERT INTO categories (name) VALUES (?)", name)
	if err != nil {
		log.Fatalf("Failed to insert category: %v", err)
		http.Error(w, "Failed to insert category", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Category created successfully"))
}

// HandleCategoryUpdate handles updating an existing category.
func HandleCategoryUpdate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	id := r.FormValue("id")
	name := r.FormValue("name")

	_, err := db.Exec("UPDATE categories SET name = ? WHERE id = ?", name, id)
	if err != nil {
		log.Fatalf("Failed to update category: %v", err)
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Category updated successfully"))
}

// HandleCategoryDelete handles deleting an existing category.
func HandleCategoryDelete(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	id := r.FormValue("id")

	_, err := db.Exec("DELETE FROM categories WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Failed to delete category: %v", err)
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Category deleted successfully"))
}
