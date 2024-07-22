// backend/controllers/menu.go
package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"golang-boilerplate/common/models/v1"
	"golang-boilerplate/config"
)

func HandleMenuIndex(w http.ResponseWriter, r *http.Request) {
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
		"backend/views/menus/index.html",
		"backend/views/menus/modals.html",
	))

	// Data for main layout template
	data := struct {
		Title      string
		MenuItems  []models.Menu
		CurrentURL string
	}{
		Title:      "Menus",
		MenuItems:  menuItems,
		CurrentURL: r.URL.Path,
	}

	// Execute main layout template with data
	if err := tmpl.ExecuteTemplate(w, "main.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandleMenuList handles listing menus.
func HandleMenuList(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Initialize a slice to hold the menu data
	var menus []models.Menu

	// Use GORM to find all menu records, ordered by ID in descending order
	if err := db.Order("id DESC").Find(&menus).Error; err != nil {
		log.Printf("Failed to query menus: %v", err)
		http.Error(w, "Failed to query menus", http.StatusInternalServerError)
		return
	}

	// Encode the menus to JSON and write to the response
	if err := json.NewEncoder(w).Encode(menus); err != nil {
		log.Printf("Failed to encode menus to JSON: %v", err)
		http.Error(w, "Failed to encode menus to JSON", http.StatusInternalServerError)
		return
	}
}

// HandleMenuCreate handles creating a new menu.
func HandleMenuCreate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Parse form values
	name := r.FormValue("name")
	link := r.FormValue("link")
	parentStr := r.FormValue("parent")

	// Convert parent to integer
	parent, err := strconv.Atoi(parentStr)
	if err != nil {
		log.Printf("Invalid parent value: %v", err)
		http.Error(w, "Invalid parent value", http.StatusBadRequest)
		return
	}

	// Create a new menu instance
	menu := models.Menu{
		Name:   name,
		Link:   link,
		Parent: parent,
	}

	// Insert the new menu using GORM
	if err := db.Create(&menu).Error; err != nil {
		log.Printf("Failed to insert menu: %v", err)
		http.Error(w, "Failed to insert menu", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Menu created successfully"))
}

// HandleMenuUpdate handles updating an existing menu.
func HandleMenuUpdate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Parse form values
	idStr := r.FormValue("id")
	name := r.FormValue("name")
	link := r.FormValue("link")
	parentStr := r.FormValue("parent")

	// Convert id and parent to integers
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid id value: %v", err)
		http.Error(w, "Invalid id value", http.StatusBadRequest)
		return
	}

	parent, err := strconv.Atoi(parentStr)
	if err != nil {
		log.Printf("Invalid parent value: %v", err)
		http.Error(w, "Invalid parent value", http.StatusBadRequest)
		return
	}

	// Find the menu by ID
	var menu models.Menu
	if err := db.First(&menu, id).Error; err != nil {
		log.Printf("Menu not found: %v", err)
		http.Error(w, "Menu not found", http.StatusNotFound)
		return
	}

	// Update the menu fields
	menu.Name = name
	menu.Link = link
	menu.Parent = parent

	// Save the updated menu
	if err := db.Save(&menu).Error; err != nil {
		log.Printf("Failed to update menu: %v", err)
		http.Error(w, "Failed to update menu", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Menu updated successfully"))
}

// HandleMenuDelete handles deleting an existing menu.
func HandleMenuDelete(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Parse form value
	idStr := r.FormValue("id")

	// Convert id to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid id value: %v", err)
		http.Error(w, "Invalid id value", http.StatusBadRequest)
		return
	}

	// Find the menu by ID
	var menu models.Menu
	if err := db.First(&menu, id).Error; err != nil {
		log.Printf("Menu not found: %v", err)
		http.Error(w, "Menu not found", http.StatusNotFound)
		return
	}

	// Delete the menu
	if err := db.Delete(&menu).Error; err != nil {
		log.Printf("Failed to delete menu: %v", err)
		http.Error(w, "Failed to delete menu", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Menu deleted successfully"))
}
