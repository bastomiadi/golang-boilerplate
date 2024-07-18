// seeder/20240717_seed_roles_table.go
package seeder

import (
	"database/sql"
	"log"
)

func SeedRolesTable(db *sql.DB) {
	query := `
    INSERT INTO roles (name) VALUES
    ('Admin'),
    ('User');`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not seed roles table: %v", err)
	}
	log.Println("Seeded roles table")
}
