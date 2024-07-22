// backend/controllers/product.go
package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"golang-boilerplate/common/models/v1"
	"golang-boilerplate/config"
)

// HandleCategoryIndex serves the index.html file for category management
func HandleProductIndex(w http.ResponseWriter, r *http.Request) {

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
		"backend/views/products/index.html",
		"backend/views/products/modals.html",
	))

	// Data for main layout template
	data := struct {
		Title      string
		MenuItems  []models.Menu
		CurrentURL string
	}{
		Title:      "Products",
		MenuItems:  menuItems,
		CurrentURL: r.URL.Path,
	}

	// Execute main layout template with data
	if err := tmpl.ExecuteTemplate(w, "main.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandleProductList handles listing products.
func HandleProductList(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	var products []models.Product

	// Fetch all products ordered by ID in descending order
	if err := db.Order("id desc").Find(&products).Error; err != nil {
		http.Error(w, "Failed to query products", http.StatusInternalServerError)
		return
	}

	// Encode the products to JSON and write the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Failed to encode products", http.StatusInternalServerError)
		return
	}
}

// HandleProductCreate handles creating a new product.
func HandleProductCreate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Parse form values
	name := r.FormValue("name")
	priceStr := r.FormValue("price")
	category := r.FormValue("category")

	// Convert price to float64
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Invalid price format", http.StatusBadRequest)
		return
	}

	// Create a new product instance
	product := models.Product{
		Name:     name,
		Price:    price,
		Category: category,
	}

	// Insert the new product into the database
	if err := db.Create(&product).Error; err != nil {
		http.Error(w, "Failed to insert product", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product created successfully"))
}

// HandleProductUpdate handles updating an existing product.
func HandleProductUpdate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Parse form values
	idStr := r.FormValue("id")
	name := r.FormValue("name")
	priceStr := r.FormValue("price")
	category := r.FormValue("category")

	// Convert id to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID format", http.StatusBadRequest)
		return
	}

	// Convert price to float64
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Invalid price format", http.StatusBadRequest)
		return
	}

	// Create a product instance with the new values
	product := models.Product{
		ID:       id,
		Name:     name,
		Price:    price,
		Category: category,
	}

	// Update the product in the database
	if err := db.Model(&product).Updates(product).Error; err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product updated successfully"))
}

// HandleProductDelete handles deleting an existing product.
func HandleProductDelete(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Get the product ID from the request
	idStr := r.FormValue("id")

	// Convert the ID from string to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID format", http.StatusBadRequest)
		return
	}

	// Create an instance of the product with the given ID
	product := models.Product{ID: id}

	// Delete the product from the database
	if err := db.Delete(&product).Error; err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product deleted successfully"))
}
