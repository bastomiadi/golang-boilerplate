// migration/20240717_create_menus_table.go
package migration

import (
	models "golang-boilerplate/common/models/v1"
	"log"

	"gorm.io/gorm"
)

func CreateMenusTable(db *gorm.DB) {
	// Drop the table if it exists
	if err := db.Migrator().DropTable(&models.Menu{}); err != nil {
		log.Fatalf("Could not drop menus table: %v", err)
	}
	log.Println("Dropped menus table if it existed")

	// Create the table
	if err := db.AutoMigrate(&models.Menu{}); err != nil {
		log.Fatalf("Could not create menus table: %v", err)
	}
	log.Println("Created menus table")
}
