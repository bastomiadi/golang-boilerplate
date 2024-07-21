package main

import (
	"golang-boilerplate/config"
	"golang-boilerplate/migration"
	"golang-boilerplate/routes"
	"golang-boilerplate/seeder"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	config.Connect()
	db := config.GetDB()
	defer db.Close()

	// migration
	migration.Migrate()

	// seeder
	seeder.Seed()

	// Initialize mux router
	router := mux.NewRouter().StrictSlash(true)
	// Register API routes
	routes.RegisterApiRoutes(router)
	// Register backend and frontend routes
	routes.RegisterBackendRoutes(router)
	routes.RegisterFrontendRoutes(router)

	// Serve static files
	router.PathPrefix("/adminlte/").Handler(http.StripPrefix("/adminlte/", http.FileServer(http.Dir("public/adminlte"))))

	// Start server
	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
