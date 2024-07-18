// seeder/20240717_seed_products_table.go
package seeder

import (
	"database/sql"
	"log"
)

func SeedProductsTable(db *sql.DB) {
	query := `
    INSERT INTO products (name, price, category) VALUES
    ('Product 1', 99.99, 'Category 1'),
    ('Product 2', 49.99, 'Category 2'),
    ('Product 3', 149.99, 'Category 1');`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not seed products table: %v", err)
	}
	log.Println("Seeded products table")
}
