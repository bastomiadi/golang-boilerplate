// migration/20240717_create_categories_table.go
package migration

import (
	"log"

	models "golang-boilerplate/common/models/v1"

	"gorm.io/gorm"
)

func CreateCategoriesTable(db *gorm.DB) {
	// Drop the table if it exists
	if err := db.Migrator().DropTable(&models.Category{}); err != nil {
		log.Fatalf("Could not drop categories table: %v", err)
	}
	log.Println("Dropped categories table if it existed")

	// Create the table
	if err := db.AutoMigrate(&models.Category{}); err != nil {
		log.Fatalf("Could not create categories table: %v", err)
	}
	log.Println("Created categories table")
}
