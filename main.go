package main

import (
	"database/sql"
	apiV1Controllers "golang-boilerplate/api/v1/controllers"
	apiV2Controllers "golang-boilerplate/api/v2/controllers"
	backendControllers "golang-boilerplate/backend/controllers"
	frontendControllers "golang-boilerplate/frontend/controllers"
	"golang-boilerplate/migration"
	"golang-boilerplate/seeder"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "root:@/golang-boilerplate")
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	// Run migrations
	migration.CreateCategoriesTable(db)
	migration.CreateProductsTable(db)

	// Run seeders
	seeder.SeedCategoriesTable(db)
	seeder.SeedProductsTable(db)

	// Serve static files
	http.Handle("/adminlte/", http.StripPrefix("/adminlte/", http.FileServer(http.Dir("public/adminlte"))))

	http.HandleFunc("/backend/dashboard", backendControllers.Dashboard)

	// Initialize routes
	http.HandleFunc("/backend/login", backendControllers.HandleLogin)
	http.HandleFunc("/backend/register", backendControllers.HandleRegister)
	http.HandleFunc("/frontend/", frontendControllers.HandleRequest)
	http.HandleFunc("/api/v1/", apiV1Controllers.HandleRequest)
	http.HandleFunc("/api/v2/", apiV2Controllers.HandleRequest)

	// Initialize routes categories
	http.HandleFunc("/backend/categories", backendControllers.HandleCategoryIndex)
	http.HandleFunc("/backend/categories/list", backendControllers.HandleCategoryList)
	http.HandleFunc("/backend/categories/create", backendControllers.HandleCategoryCreate)
	http.HandleFunc("/backend/categories/update", backendControllers.HandleCategoryUpdate)
	http.HandleFunc("/backend/categories/delete", backendControllers.HandleCategoryDelete)

	// Initializa routes products
	http.HandleFunc("/backend/products", backendControllers.HandleProductIndex)
	http.HandleFunc("/backend/products/list", backendControllers.HandleProductList)
	http.HandleFunc("/backend/products/create", backendControllers.HandleProductCreate)
	http.HandleFunc("/backend/products/update", backendControllers.HandleProductUpdate)
	http.HandleFunc("/backend/products/delete", backendControllers.HandleProductDelete)

	// Start server
	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
