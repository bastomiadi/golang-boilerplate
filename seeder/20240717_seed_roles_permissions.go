// seeder/20240717_seed_permissions_table.go
package seeder

import (
	"database/sql"
	"log"
)

func SeedRolePermissionsTable(db *sql.DB) {
	query := `
    INSERT INTO role_permissions (role_id, permission_id) VALUES
    (1, 1),
    (2, 2);`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not seed role permissions table: %v", err)
	}
	log.Println("Seeded role permissions table")
}
