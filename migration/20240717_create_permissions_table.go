// migration/20240717_update_permissions_table.go
package migration

import (
	models "golang-boilerplate/common/models/v1"
	"log"

	"gorm.io/gorm"
)

func CreatePermissionsTable(db *gorm.DB) {

	// Drop the table if it exists
	if err := db.Migrator().DropTable(&models.Permission{}); err != nil {
		log.Fatalf("Could not drop permissions table: %v", err)
	}
	log.Println("Dropped permissions table if it existed")

	// Perform the migration
	if err := db.AutoMigrate(&models.Permission{}); err != nil {
		log.Fatalf("Could not migrate permissions table: %v", err)
	}
	log.Println("Migrated permissions table")
}
