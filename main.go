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

	migration.Migrate()

	seeder.Seed()

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
