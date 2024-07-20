// seeder/20240717_seed_permissions_table.go
package seeder

import (
	"database/sql"
	"log"
)

func SeedUserRolesTable(db *sql.DB) {
	query := `
    INSERT INTO user_roles (user_id, role_id) VALUES
    (1, 1),
    (2, 2);`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not seed user roles table: %v", err)
	}
	log.Println("Seeded user roles table")
}
