// backend/controllers/permission.go
package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"golang-boilerplate/backend/models"
	"golang-boilerplate/config"
)

// HandlePermissionList handles listing permissions.
func HandlePermissionList(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	permissions := []models.Permission{}

	rows, err := db.Query("SELECT id, name FROM permissions ORDER BY id DESC")
	if err != nil {
		log.Fatalf("Failed to query permissions: %v", err)
		http.Error(w, "Failed to query permissions", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var permission models.Permission
		if err := rows.Scan(&permission.ID, &permission.Name); err != nil {
			log.Fatalf("Failed to scan permission: %v", err)
			http.Error(w, "Failed to scan permissions", http.StatusInternalServerError)
			return
		}
		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
		http.Error(w, "Error iterating over permissions", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(permissions)
}

// HandlePermissionCreate handles creating a new permission.
func HandlePermissionCreate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	name := r.FormValue("name")

	_, err := db.Exec("INSERT INTO permissions (name) VALUES (?)", name)
	if err != nil {
		log.Fatalf("Failed to insert permission: %v", err)
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

	_, err := db.Exec("UPDATE permissions SET name = ? WHERE id = ?", name, id)
	if err != nil {
		log.Fatalf("Failed to update permission: %v", err)
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

	_, err := db.Exec("DELETE FROM permissions WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Failed to delete permission: %v", err)
		http.Error(w, "Failed to delete permission", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Permission deleted successfully"))
}
