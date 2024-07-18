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
	migration.CreateUsersTable(db)
	migration.CreateMenusTable(db)
	migration.UpdateRolesTable(db)           // Add this line
	migration.UpdatePermissionsTable(db)     // Add this line
	migration.CreateRolePermissionsTable(db) // Add this line
	migration.CreateUserRolesTable(db)       // Add this line

	// Run seeders
	seeder.SeedCategoriesTable(db)
	seeder.SeedProductsTable(db)
	seeder.SeedUsersTable(db)
	seeder.SeedMenusTable(db)
	seeder.SeedRolesTable(db)       // Add this line
	seeder.SeedPermissionsTable(db) // Add this line

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

	http.HandleFunc("/backend/users/list", backendControllers.HandleUserList)
	http.HandleFunc("/backend/users/create", backendControllers.HandleUserCreate)
	http.HandleFunc("/backend/users/update", backendControllers.HandleUserUpdate)
	http.HandleFunc("/backend/users/delete", backendControllers.HandleUserDelete)

	http.HandleFunc("/backend/menus/list", backendControllers.HandleMenuList)
	http.HandleFunc("/backend/menus/create", backendControllers.HandleMenuCreate)
	http.HandleFunc("/backend/menus/update", backendControllers.HandleMenuUpdate)
	http.HandleFunc("/backend/menus/delete", backendControllers.HandleMenuDelete)

	http.HandleFunc("/backend/roles/list", backendControllers.HandleRoleList)     // Add this line
	http.HandleFunc("/backend/roles/create", backendControllers.HandleRoleCreate) // Add this line
	http.HandleFunc("/backend/roles/update", backendControllers.HandleRoleUpdate) // Add this line
	http.HandleFunc("/backend/roles/delete", backendControllers.HandleRoleDelete) // Add this line

	http.HandleFunc("/backend/permissions/list", backendControllers.HandlePermissionList)     // Add this line
	http.HandleFunc("/backend/permissions/create", backendControllers.HandlePermissionCreate) // Add this line
	http.HandleFunc("/backend/permissions/update", backendControllers.HandlePermissionUpdate) // Add this line
	http.HandleFunc("/backend/permissions/delete", backendControllers.HandlePermissionDelete) // Add this line

	// Start server
	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
