package controllers

import (
	"encoding/json"
	models "golang-boilerplate/common/models/v1"
	"golang-boilerplate/config"
	"net/http"
	"strconv"
	"time"
)

// Get all categories
func GetCategories(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	var categories []models.Category
	// Query categories using GORM
	if err := db.Select("id, name, created_at").Find(&categories).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set header and encode response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Create new category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Set created_at to current time
	category.CreatedAt = time.Now()

	// Create a new category record using GORM
	if err := db.Create(&category).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set header and encode response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Update category by ID
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Update the category record using GORM
	if err := db.Model(&models.Category{}).Where("id = ?", category.ID).Updates(map[string]interface{}{
		"name": category.Name,
	}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete category by ID
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Only allow DELETE method
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve category ID from query parameter
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	// Execute DELETE query in the database using GORM
	if err := db.Where("id = ?", id).Delete(&models.Category{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
}
