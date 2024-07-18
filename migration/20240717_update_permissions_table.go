// migration/20240717_update_permissions_table.go
package migration

import (
	"database/sql"
	"log"
)

func UpdatePermissionsTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS permissions (
        id INT AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        PRIMARY KEY (id)
    );`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not create permissions table: %v", err)
	}
	log.Println("Created permissions table")
}
