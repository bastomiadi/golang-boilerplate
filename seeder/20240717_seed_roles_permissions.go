// seeder/20240717_seed_permissions_table.go
package seeder

import (
	"golang-boilerplate/common/models/v1"
	"golang-boilerplate/config"
	"log"
)

func SeedRolePermissionsTable() {
	db := config.GetDB()

	// Define initial role permissions
	rolePermissions := []models.RolePermission{
		{RoleID: 1, PermissionID: 1},
		{RoleID: 2, PermissionID: 2},
	}

	// Perform the seed operation
	if err := db.Create(&rolePermissions).Error; err != nil {
		log.Fatalf("Could not seed role_permissions table: %v", err)
	}

	log.Println("Seeded role_permissions table")
}
