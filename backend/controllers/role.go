// backend/controllers/role.go
package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"golang-boilerplate/backend/models"
)

// HandleRoleList handles listing roles.
func HandleRoleList(w http.ResponseWriter, r *http.Request) {
	db := getDB() // Function to get DB connection
	defer db.Close()

	roles := []models.Role{}

	rows, err := db.Query("SELECT id, name FROM roles ORDER BY id DESC")
	if err != nil {
		log.Fatalf("Failed to query roles: %v", err)
		http.Error(w, "Failed to query roles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var role models.Role
		if err := rows.Scan(&role.ID, &role.Name); err != nil {
			log.Fatalf("Failed to scan role: %v", err)
			http.Error(w, "Failed to scan roles", http.StatusInternalServerError)
			return
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
		http.Error(w, "Error iterating over roles", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(roles)
}

// HandleRoleCreate handles creating a new role.
func HandleRoleCreate(w http.ResponseWriter, r *http.Request) {
	db := getDB() // Function to get DB connection
	defer db.Close()

	name := r.FormValue("name")

	_, err := db.Exec("INSERT INTO roles (name) VALUES (?)", name)
	if err != nil {
		log.Fatalf("Failed to insert role: %v", err)
		http.Error(w, "Failed to insert role", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Role created successfully"))
}

// HandleRoleUpdate handles updating an existing role.
func HandleRoleUpdate(w http.ResponseWriter, r *http.Request) {
	db := getDB() // Function to get DB connection
	defer db.Close()

	id := r.FormValue("id")
	name := r.FormValue("name")

	_, err := db.Exec("UPDATE roles SET name = ? WHERE id = ?", name, id)
	if err != nil {
		log.Fatalf("Failed to update role: %v", err)
		http.Error(w, "Failed to update role", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Role updated successfully"))
}

// HandleRoleDelete handles deleting an existing role.
func HandleRoleDelete(w http.ResponseWriter, r *http.Request) {
	db := getDB() // Function to get DB connection
	defer db.Close()

	id := r.FormValue("id")

	_, err := db.Exec("DELETE FROM roles WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Failed to delete role: %v", err)
		http.Error(w, "Failed to delete role", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Role deleted successfully"))
}
