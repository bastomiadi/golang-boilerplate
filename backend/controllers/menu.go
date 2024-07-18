// backend/controllers/menu.go
package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"golang-boilerplate/backend/models"
	"golang-boilerplate/config"
)

// HandleMenuList handles listing menus.
func HandleMenuList(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	menus := []models.Menu{}

	rows, err := db.Query("SELECT id, name, link, parent FROM menus ORDER BY id DESC")
	if err != nil {
		log.Fatalf("Failed to query menus: %v", err)
		http.Error(w, "Failed to query menus", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var menu models.Menu
		if err := rows.Scan(&menu.ID, &menu.Name, &menu.Link, &menu.Parent); err != nil {
			log.Fatalf("Failed to scan menu: %v", err)
			http.Error(w, "Failed to scan menus", http.StatusInternalServerError)
			return
		}
		menus = append(menus, menu)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
		http.Error(w, "Error iterating over menus", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(menus)
}

// HandleMenuCreate handles creating a new menu.
func HandleMenuCreate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	name := r.FormValue("name")
	link := r.FormValue("link")
	parent := r.FormValue("parent")

	_, err := db.Exec("INSERT INTO menus (name, link, parent) VALUES (?, ?, ?)", name, link, parent)
	if err != nil {
		log.Fatalf("Failed to insert menu: %v", err)
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

	id := r.FormValue("id")
	name := r.FormValue("name")
	link := r.FormValue("link")
	parent := r.FormValue("parent")

	_, err := db.Exec("UPDATE menus SET name = ?, link = ?, parent = ? WHERE id = ?", name, link, parent, id)
	if err != nil {
		log.Fatalf("Failed to update menu: %v", err)
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

	id := r.FormValue("id")

	_, err := db.Exec("DELETE FROM menus WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Failed to delete menu: %v", err)
		http.Error(w, "Failed to delete menu", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Menu deleted successfully"))
}
