// migration/20240717_create_user_roles_table.go
package migration

import (
	models "golang-boilerplate/common/models/v1"
	"log"

	"gorm.io/gorm"
)

func CreateUserRolesTable(db *gorm.DB) {
	// Drop the table if it exists
	if err := db.Migrator().DropTable(&models.UserRole{}); err != nil {
		log.Fatalf("Could not drop user_roles table: %v", err)
	}
	log.Println("Dropped user_roles table if it existed")

	// Create the table
	if err := db.AutoMigrate(&models.UserRole{}); err != nil {
		log.Fatalf("Could not create user_roles table: %v", err)
	}
	log.Println("Created user_roles table")
}
