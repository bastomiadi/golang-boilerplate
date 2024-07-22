// seeder/20240717_seed_permissions_table.go
package seeder

import (
	"golang-boilerplate/common/models/v1"
	"golang-boilerplate/config"
	"log"
)

func SeedUserRolesTable() {
	db := config.GetDB()

	// Define initial user roles
	userRoles := []models.UserRole{
		{UserID: 1, RoleID: 1},
		{UserID: 2, RoleID: 2},
	}

	// Perform the seed operation
	if err := db.Create(&userRoles).Error; err != nil {
		log.Fatalf("Could not seed user roles table: %v", err)
	}

	log.Println("Seeded user roles table")
}
