// seeder/20230717_seed_categories_table.go
package seeder

import (
	"database/sql"
	"log"
)

func SeedCategoriesTable(db *sql.DB) {
	query := `
    INSERT INTO categories (name) VALUES
    ('Category 1'),
    ('Category 2'),
    ('Category 3');`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not seed categories table: %v", err)
	}
	log.Println("Seeded categories table")
}
