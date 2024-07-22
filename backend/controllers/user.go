// backend/controllers/user.go
package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"golang-boilerplate/backend/models"
	"golang-boilerplate/config"
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

	users := []models.User{}

	rows, err := db.Query("SELECT id, name, username, email FROM users ORDER BY id DESC")
	if err != nil {
		log.Fatalf("Failed to query users: %v", err)
		http.Error(w, "Failed to query users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email); err != nil {
			log.Fatalf("Failed to scan user: %v", err)
			http.Error(w, "Failed to scan users", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
		http.Error(w, "Error iterating over users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// HandleUserCreate handles creating a new user.
func HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	name := r.FormValue("name")
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	_, err := db.Exec("INSERT INTO users (name, username, email, password) VALUES (?, ?, ?, ?)", name, username, email, password)
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
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

	id := r.FormValue("id")
	name := r.FormValue("name")
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	_, err := db.Exec("UPDATE users SET name = ?, username = ?, email = ?, password = ? WHERE id = ?", name, username, email, password, id)
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
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

	id := r.FormValue("id")

	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Failed to delete user: %v", err)
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
