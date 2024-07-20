// seeder/20240717_seed_menus_table.go
package seeder

import (
	"database/sql"
	"log"
)

func SeedMenusTable(db *sql.DB) {
	query := `
    INSERT INTO menus (name, link, parent, icon, display_order) VALUES
    ('Dashboard', '/backend/dashboard', 0, '', 0),
    ('Master', '', 0, '', 0),
    ('Categories', '/backend/categories', 2, '', 0),
    ('Products', '/backend/products', 2, '', 0),
	('RBAC', '', 0, '', 0),
	('Users', '/backend/users', 5, '', 0),
	('Roles', '/backend/roles', 5, '', 0),
	('Permissions', '/backend/permissions', 5, '', 0),
	('Menus', '/backend/menus', 5, '', 0);`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not seed menus table: %v", err)
	}
	log.Println("Seeded menus table")
}
