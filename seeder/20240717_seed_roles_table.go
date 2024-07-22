// seeder/20240717_seed_roles_table.go
package seeder

import (
	"golang-boilerplate/common/models/v1"
	"golang-boilerplate/config"
	"log"
)

func SeedRolesTable() {
	db := config.GetDB()

	// Define initial roles
	roles := []models.Role{
		{Name: "Admin"},
		{Name: "User"},
	}

	// Perform the seed operation
	if err := db.Create(&roles).Error; err != nil {
		log.Fatalf("Could not seed roles table: %v", err)
	}

	log.Println("Seeded roles table")
}
