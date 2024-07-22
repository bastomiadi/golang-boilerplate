// backend/controllers/user.go
package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"golang-boilerplate/common/models/v1"
	"golang-boilerplate/config"
	"golang-boilerplate/utils"
)

func HandleUserIndex(w http.ResponseWriter, r *http.Request) {

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
		"backend/views/users/index.html",
		"backend/views/users/modals.html",
	))

	// Data for main layout template
	data := struct {
		Title      string
		MenuItems  []models.Menu
		CurrentURL string
	}{
		Title:      "Users",
		MenuItems:  menuItems,
		CurrentURL: r.URL.Path,
	}

	// Execute main layout template with data
	if err := tmpl.ExecuteTemplate(w, "main.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandleUserList handles listing users.
func HandleUserList(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	var users []models.User

	// Fetch users from the database using GORM
	if err := db.Order("id desc").Find(&users).Error; err != nil {
		http.Error(w, "Failed to query users", http.StatusInternalServerError)
		return
	}

	// Return users as JSON response
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
	}
}

// HandleUserCreate handles creating a new user.
func HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	name := r.FormValue("name")
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Hash the password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Create a new user instance
	user := models.User{
		Name:     name,
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

	// Insert user into the database using GORM
	if err := db.Create(&user).Error; err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created successfully"))
}

// HandleUserUpdate handles updating an existing user.
func HandleUserUpdate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Parse form values
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")
	name := r.FormValue("name")
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Create a map of fields to update
	updateFields := map[string]interface{}{
		"name":     name,
		"username": username,
		"email":    email,
	}

	// If a password is provided, hash it and include it in the update
	if password != "" {
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		updateFields["password"] = hashedPassword
	}

	// Perform the update using GORM
	if err := db.Model(&models.User{}).Where("id = ?", id).Updates(updateFields).Error; err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}

// HandleUserDelete handles deleting an existing user.
func HandleUserDelete(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Parse form values
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")

	// Create a new User instance for GORM
	var user models.User

	// Perform the delete operation using GORM
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		log.Printf("Failed to delete user: %v", err)
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
