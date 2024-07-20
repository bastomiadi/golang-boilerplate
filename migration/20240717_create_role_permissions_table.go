// migration/20240717_create_role_permissions_table.go
package migration

import (
	"database/sql"
	"log"
)

func CreateRolePermissionsTable(db *sql.DB) {
	dropQuery := `
    DROP TABLE IF EXISTS role_permissions;
    `
	createQuery := `
    CREATE TABLE IF NOT EXISTS role_permissions (
        role_id INT,
        permission_id INT,
        FOREIGN KEY (role_id) REFERENCES roles(id),
        FOREIGN KEY (permission_id) REFERENCES permissions(id),
        PRIMARY KEY (role_id, permission_id)
    );`

	// Execute the drop query
	if _, err := db.Exec(dropQuery); err != nil {
		log.Fatalf("Could not drop role_permissions table: %v", err)
	}
	log.Println("Dropped role_permissions table if it existed")

	// Execute the create query
	if _, err := db.Exec(createQuery); err != nil {
		log.Fatalf("Could not create role_permissions table: %v", err)
	}
	log.Println("Created role_permissions table")
}
