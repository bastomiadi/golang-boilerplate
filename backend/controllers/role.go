// backend/controllers/role.go
package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"golang-boilerplate/common/models/v1"
	"golang-boilerplate/config"
)

func HandleRoleIndex(w http.ResponseWriter, r *http.Request) {

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
		"backend/views/roles/index.html",
		"backend/views/roles/modals.html",
	))

	// Data for main layout template
	data := struct {
		Title      string
		MenuItems  []models.Menu
		CurrentURL string
	}{
		Title:      "Roles",
		MenuItems:  menuItems,
		CurrentURL: r.URL.Path,
	}

	// Execute main layout template with data
	if err := tmpl.ExecuteTemplate(w, "main.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandleRoleList handles listing roles.
func HandleRoleList(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Define a slice to hold roles
	roles := []models.Role{}

	// Query roles from the database
	if err := db.Order("id desc").Find(&roles).Error; err != nil {
		http.Error(w, "Failed to query roles", http.StatusInternalServerError)
		return
	}

	// Encode roles as JSON and send the response
	if err := json.NewEncoder(w).Encode(roles); err != nil {
		http.Error(w, "Failed to encode roles", http.StatusInternalServerError)
		return
	}
}

// HandleRoleCreate handles creating a new role.
func HandleRoleCreate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Retrieve form value for the role name
	name := r.FormValue("name")

	// Create a new role instance
	role := models.Role{Name: name}

	// Insert the new role into the database
	if err := db.Create(&role).Error; err != nil {
		http.Error(w, "Failed to insert role", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Role created successfully"))
}

// HandleRoleUpdate handles updating an existing role.
func HandleRoleUpdate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Retrieve form values for the role ID and name
	id := r.FormValue("id")
	name := r.FormValue("name")

	// Find the role by ID
	var role models.Role
	if err := db.First(&role, id).Error; err != nil {
		http.Error(w, "Role not found", http.StatusNotFound)
		return
	}

	// Update the role's name
	role.Name = name
	if err := db.Save(&role).Error; err != nil {
		http.Error(w, "Failed to update role", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Role updated successfully"))
}

// HandleRoleDelete handles deleting an existing role.
func HandleRoleDelete(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Retrieve form value for the role ID
	id := r.FormValue("id")

	// Delete the role by ID
	if err := db.Delete(&models.Role{}, id).Error; err != nil {
		http.Error(w, "Failed to delete role", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Role deleted successfully"))
}
