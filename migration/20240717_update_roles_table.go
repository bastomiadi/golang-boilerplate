// migration/20240717_update_roles_table.go
package migration

import (
	"database/sql"
	"log"
)

func UpdateRolesTable(db *sql.DB) {

	dropRolesPermissionsQuery := `
    DROP TABLE IF EXISTS role_permissions;
    `

	// Execute the drop query
	if _, err := db.Exec(dropRolesPermissionsQuery); err != nil {
		log.Fatalf("Could not drop roles permissions table: %v", err)
	}
	log.Println("Dropped roles permissions table if it existed")

	dropQuery := `
    DROP TABLE IF EXISTS roles;
    `
	createQuery := `
    CREATE TABLE IF NOT EXISTS roles (
        id INT AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        PRIMARY KEY (id)
    );`

	// Execute the drop query
	if _, err := db.Exec(dropQuery); err != nil {
		log.Fatalf("Could not drop roles table: %v", err)
	}
	log.Println("Dropped roles table if it existed")

	// Execute the create query
	if _, err := db.Exec(createQuery); err != nil {
		log.Fatalf("Could not create roles table: %v", err)
	}
	log.Println("Created roles table")
}
