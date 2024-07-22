// backend/controllers/permission.go
package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"golang-boilerplate/common/models/v1"
	"golang-boilerplate/config"
)

func HandlePermissionIndex(w http.ResponseWriter, r *http.Request) {

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
		"backend/views/permissions/index.html",
		"backend/views/permissions/modals.html",
	))

	// Data for main layout template
	data := struct {
		Title      string
		MenuItems  []models.Menu
		CurrentURL string
	}{
		Title:      "Permissions",
		MenuItems:  menuItems,
		CurrentURL: r.URL.Path,
	}

	// Execute main layout template with data
	if err := tmpl.ExecuteTemplate(w, "main.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandlePermissionList handles listing permissions.
func HandlePermissionList(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	var permissions []models.Permission

	// Use GORM to find all permissions, ordered by ID descending
	if err := db.Order("id desc").Find(&permissions).Error; err != nil {
		log.Printf("Failed to query permissions: %v", err)
		http.Error(w, "Failed to query permissions", http.StatusInternalServerError)
		return
	}

	// Encode the result to JSON and send it back
	if err := json.NewEncoder(w).Encode(permissions); err != nil {
		log.Printf("Failed to encode permissions to JSON: %v", err)
		http.Error(w, "Failed to encode permissions", http.StatusInternalServerError)
	}
}

// HandlePermissionCreate handles creating a new permission.
func HandlePermissionCreate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	name := r.FormValue("name")

	// Create a new Permission instance
	permission := models.Permission{Name: name}

	// Use GORM to insert the new permission record
	if err := db.Create(&permission).Error; err != nil {
		http.Error(w, "Failed to insert permission", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Permission created successfully"))
}

// HandlePermissionUpdate handles updating an existing permission.
func HandlePermissionUpdate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	id := r.FormValue("id")
	name := r.FormValue("name")

	// Create a new Permission instance with the provided ID
	var permission models.Permission

	// Find the permission by ID
	if err := db.First(&permission, id).Error; err != nil {
		http.Error(w, "Permission not found", http.StatusNotFound)
		return
	}

	// Update the permission's name
	permission.Name = name

	// Save the updated permission record
	if err := db.Save(&permission).Error; err != nil {
		http.Error(w, "Failed to update permission", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Permission updated successfully"))
}

// HandlePermissionDelete handles deleting an existing permission.
func HandlePermissionDelete(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	id := r.FormValue("id")

	// Create a new Permission instance with the provided ID
	var permission models.Permission

	// Find the permission by ID
	if err := db.First(&permission, id).Error; err != nil {
		http.Error(w, "Permission not found", http.StatusNotFound)
		return
	}

	// Delete the permission
	if err := db.Delete(&permission).Error; err != nil {
		http.Error(w, "Failed to delete permission", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Permission deleted successfully"))
}
