// seeder/20240717_seed_menus_table.go
package seeder

import (
	"database/sql"
	"log"
)

func SeedMenusTable(db *sql.DB) {
	query := `
    INSERT INTO menus (name, link, parent) VALUES
    ('Home', '/', 0),
    ('About', '/about', 0),
    ('Services', '/services', 0),
    ('Contact', '/contact', 0);`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not seed menus table: %v", err)
	}
	log.Println("Seeded menus table")
}
