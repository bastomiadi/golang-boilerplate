// migration/20240717_create_products_table.go
package migration

import (
	"log"

	models "golang-boilerplate/common/models/v1"

	"gorm.io/gorm"
)

func CreateProductsTable(db *gorm.DB) {
	// Drop the table if it exists
	if err := db.Migrator().DropTable(&models.Product{}); err != nil {
		log.Fatalf("Could not drop products table: %v", err)
	}
	log.Println("Dropped products table if it existed")

	// Create the table
	if err := db.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("Could not create products table: %v", err)
	}
	log.Println("Created products table")
}
