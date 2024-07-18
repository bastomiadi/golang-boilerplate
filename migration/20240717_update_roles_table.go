// migration/20240717_update_roles_table.go
package migration

import (
	"database/sql"
	"log"
)

func UpdateRolesTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS roles (
        id INT AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        PRIMARY KEY (id)
    );`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not create roles table: %v", err)
	}
	log.Println("Created roles table")
}
