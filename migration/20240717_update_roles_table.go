// migration/20240717_update_roles_table.go
package migration

import (
	"log"

	models "golang-boilerplate/common/models/v1"

	"gorm.io/gorm"
)

func UpdateRolesTable(db *gorm.DB) {

	// Drop the table if it exists
	if err := db.Migrator().DropTable(&models.Role{}); err != nil {
		log.Fatalf("Could not drop roles table: %v", err)
	}
	log.Println("Dropped roles table if it existed")

	// Create the table
	if err := db.AutoMigrate(&models.Role{}); err != nil {
		log.Fatalf("Could not create roles table: %v", err)
	}
	log.Println("Created roles table")
}
