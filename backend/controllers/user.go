// backend/controllers/user.go
package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"golang-boilerplate/backend/models"
)

// HandleUserList handles listing users.
func HandleUserList(w http.ResponseWriter, r *http.Request) {
	db := getDB() // Function to get DB connection
	defer db.Close()

	users := []models.User{}

	rows, err := db.Query("SELECT id, username, email, role FROM users ORDER BY id DESC")
	if err != nil {
		log.Fatalf("Failed to query users: %v", err)
		http.Error(w, "Failed to query users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
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
	db := getDB() // Function to get DB connection
	defer db.Close()

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	role := r.FormValue("role")

	_, err := db.Exec("INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)", username, email, password, role)
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
	db := getDB() // Function to get DB connection
	defer db.Close()

	id := r.FormValue("id")
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	role := r.FormValue("role")

	_, err := db.Exec("UPDATE users SET username = ?, email = ?, password = ?, role = ? WHERE id = ?", username, email, password, role, id)
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
	db := getDB() // Function to get DB connection
	defer db.Close()

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
