// migration/20240717_create_products_table.go
package migration

import (
	"database/sql"
	"log"
)

func CreateProductsTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS products (
        id INT AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        price DECIMAL(10, 2) NOT NULL,
        category VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    );`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not create products table: %v", err)
	}
	log.Println("Created products table")
}
