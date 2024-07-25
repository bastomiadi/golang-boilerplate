// migration/20240717_create_users_table.go
package migration

import (
	models "golang-boilerplate/common/models/v1"
	"log"

	"gorm.io/gorm"
)

func CreateProfilesTable(db *gorm.DB) {
	// Drop the table if it exists
	if err := db.Migrator().DropTable(&models.Profile{}); err != nil {
		log.Fatalf("Could not drop users table: %v", err)
	}
	log.Println("Dropped profiles table if it existed")

	// Create the table
	if err := db.AutoMigrate(&models.Profile{}); err != nil {
		log.Fatalf("Could not create profiles table: %v", err)
	}
	log.Println("Created profiles table")
}
