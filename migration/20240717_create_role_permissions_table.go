// migration/20240717_create_role_permissions_table.go
package migration

import (
	models "golang-boilerplate/common/models/v1"
	"log"

	"gorm.io/gorm"
)

func CreateRolePermissionsTable(db *gorm.DB) {
	// Drop the table if it exists
	if err := db.Migrator().DropTable(&models.RolePermission{}); err != nil {
		log.Fatalf("Could not drop role_permissions table: %v", err)
	}
	log.Println("Dropped role_permissions table if it existed")

	// Create the table
	if err := db.AutoMigrate(&models.RolePermission{}); err != nil {
		log.Fatalf("Could not create role_permissions table: %v", err)
	}
	log.Println("Created role_permissions table")
}
