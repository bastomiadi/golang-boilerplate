// migration/20240717_create_role_permissions_table.go
package migration

import (
	"database/sql"
	"log"
)

func CreateRolePermissionsTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS role_permissions (
        role_id INT,
        permission_id INT,
        FOREIGN KEY (role_id) REFERENCES roles(id),
        FOREIGN KEY (permission_id) REFERENCES permissions(id),
        PRIMARY KEY (role_id, permission_id)
    );`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not create role_permissions table: %v", err)
	}
	log.Println("Created role_permissions table")
}
