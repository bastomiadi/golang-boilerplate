// seeder/20240717_seed_permissions_table.go
package seeder

import (
	"database/sql"
	"log"
)

func SeedPermissionsTable(db *sql.DB) {
	query := `
    INSERT INTO permissions (name) VALUES
    ('View Users'),
    ('Create Users'),
    ('Edit Users'),
    ('Delete Users'),
    ('View Roles'),
    ('Create Roles'),
    ('Edit Roles'),
    ('Delete Roles'),
    ('View Permissions'),
    ('Create Permissions'),
    ('Edit Permissions'),
    ('Delete Permissions'),
    ('View Menus'),
    ('Create Menus'),
    ('Edit Menus'),
    ('Delete Menus');`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not seed permissions table: %v", err)
	}
	log.Println("Seeded permissions table")
}
