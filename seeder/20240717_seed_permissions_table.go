// seeder/20240717_seed_permissions_table.go
package seeder

import (
	"golang-boilerplate/common/models/v1"
	"golang-boilerplate/config"
	"log"
)

func SeedPermissionsTable() {
	db := config.GetDB()

	// Define initial permissions
	permissions := []models.Permission{
		{Name: "View Users"},
		{Name: "Create Users"},
		{Name: "Edit Users"},
		{Name: "Delete Users"},
		{Name: "View Roles"},
		{Name: "Create Roles"},
		{Name: "Edit Roles"},
		{Name: "Delete Roles"},
		{Name: "View Permissions"},
		{Name: "Create Permissions"},
		{Name: "Edit Permissions"},
		{Name: "Delete Permissions"},
		{Name: "View Menus"},
		{Name: "Create Menus"},
		{Name: "Edit Menus"},
		{Name: "Delete Menus"},
	}

	// Perform the seed operation
	if err := db.Create(&permissions).Error; err != nil {
		log.Fatalf("Could not seed permissions table: %v", err)
	}

	log.Println("Seeded permissions table")
}
