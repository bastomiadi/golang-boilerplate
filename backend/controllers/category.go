// backend/controllers/category.go
package controllers

import (
	"encoding/json"
	"golang-boilerplate/common/models/v1"
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

	var categories []models.Category

	// Query categories using GORM
	if err := db.Order("id desc").Find(&categories).Error; err != nil {
		log.Printf("Failed to query categories: %v", err)
		http.Error(w, "Failed to query categories", http.StatusInternalServerError)
		return
	}

	// Encode categories to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(categories); err != nil {
		log.Printf("Failed to encode categories to JSON: %v", err)
		http.Error(w, "Failed to encode categories", http.StatusInternalServerError)
		return
	}
}

// HandleCategoryCreate handles creating a new category.
func HandleCategoryCreate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Parse the request form values
	if err := r.ParseForm(); err != nil {
		log.Printf("Failed to parse form values: %v", err)
		http.Error(w, "Failed to parse form values", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")

	// Create a new category instance
	category := models.Category{Name: name}

	// Save the category to the database using GORM
	if err := db.Create(&category).Error; err != nil {
		log.Printf("Failed to insert category: %v", err)
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

	// Parse the request form values
	if err := r.ParseForm(); err != nil {
		log.Printf("Failed to parse form values: %v", err)
		http.Error(w, "Failed to parse form values", http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")
	name := r.FormValue("name")

	// Find the category by ID
	var category models.Category
	if err := db.First(&category, id).Error; err != nil {
		log.Printf("Failed to find category: %v", err)
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	// Update the category's name
	category.Name = name

	// Save the updated category to the database using GORM
	if err := db.Save(&category).Error; err != nil {
		log.Printf("Failed to update category: %v", err)
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

	// Parse the request form values
	if err := r.ParseForm(); err != nil {
		log.Printf("Failed to parse form values: %v", err)
		http.Error(w, "Failed to parse form values", http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")

	// Find the category by ID
	var category models.Category
	if err := db.First(&category, id).Error; err != nil {
		log.Printf("Failed to find category: %v", err)
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	// Delete the category from the database using GORM
	if err := db.Delete(&category).Error; err != nil {
		log.Printf("Failed to delete category: %v", err)
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Category deleted successfully"))
}
