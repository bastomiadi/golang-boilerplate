// backend/controllers/product.go
package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"golang-boilerplate/backend/models"
	"golang-boilerplate/config"
)

// HandleCategoryIndex serves the index.html file for category management
func HandleProductIndex(w http.ResponseWriter, r *http.Request) {
	// Parse main layout template and content template
	tmpl := template.Must(template.ParseFiles(
		"backend/views/partials/header.html",
		"backend/views/partials/sidebar.html",
		"backend/views/partials/footer.html",
		"backend/views/layouts/main.html",
		"backend/views/products/index.html",
		"backend/views/products/modals.html",
	))

	// Data for main layout template
	data := struct {
		Title string
	}{
		Title: "Products",
	}

	// Execute main layout template with data
	if err := tmpl.ExecuteTemplate(w, "main.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandleProductList handles listing products.
func HandleProductList(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	products := []models.Product{}

	rows, err := db.Query("SELECT id, name, price, category FROM products ORDER BY id DESC")
	if err != nil {
		log.Fatalf("Failed to query products: %v", err)
		http.Error(w, "Failed to query products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Category); err != nil {
			log.Fatalf("Failed to scan product: %v", err)
			http.Error(w, "Failed to scan products", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
		http.Error(w, "Error iterating over products", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

// HandleProductCreate handles creating a new product.
func HandleProductCreate(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	name := r.FormValue("name")
	price := r.FormValue("price")
	category := r.FormValue("category")

	_, err := db.Exec("INSERT INTO products (name, price, category) VALUES (?, ?, ?)", name, price, category)
	if err != nil {
		log.Fatalf("Failed to insert product: %v", err)
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

	id := r.FormValue("id")
	name := r.FormValue("name")
	price := r.FormValue("price")
	category := r.FormValue("category")

	_, err := db.Exec("UPDATE products SET name = ?, price = ?, category = ? WHERE id = ?", name, price, category, id)
	if err != nil {
		log.Fatalf("Failed to update product: %v", err)
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

	id := r.FormValue("id")

	_, err := db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Failed to delete product: %v", err)
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product deleted successfully"))
}
