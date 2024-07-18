package main

import (
	"golang-boilerplate/config"
	"golang-boilerplate/migration"
	"golang-boilerplate/routes"
	"golang-boilerplate/seeder"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database
	config.Connect()
	db := config.GetDB()
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

	// Register routes
	routes.RegisterBackendRoutes()
	routes.RegisterFrontendRoutes()
	routes.RegisterAPIRoutes()

	// Serve static files
	http.Handle("/adminlte/", http.StripPrefix("/adminlte/", http.FileServer(http.Dir("public/adminlte"))))

	// Start server
	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
