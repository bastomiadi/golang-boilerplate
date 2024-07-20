// migration/20240717_update_permissions_table.go
package migration

import (
	"database/sql"
	"log"
)

func UpdatePermissionsTable(db *sql.DB) {
	dropQuery := `
    DROP TABLE IF EXISTS permissions;
    `
	createQuery := `
    CREATE TABLE IF NOT EXISTS permissions (
        id INT AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        PRIMARY KEY (id)
    );`

	// Execute the drop query
	if _, err := db.Exec(dropQuery); err != nil {
		log.Fatalf("Could not drop permissions table: %v", err)
	}
	log.Println("Dropped permissions table if it existed")

	// Execute the create query
	if _, err := db.Exec(createQuery); err != nil {
		log.Fatalf("Could not create permissions table: %v", err)
	}
	log.Println("Created permissions table")
}
