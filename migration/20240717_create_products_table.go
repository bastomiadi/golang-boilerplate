// migration/20240717_create_products_table.go
package migration

import (
	"database/sql"
	"log"
)

func CreateProductsTable(db *sql.DB) {
	dropQuery := `
    DROP TABLE IF EXISTS products;
    `
	createQuery := `
    CREATE TABLE IF NOT EXISTS products (
        id INT AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        price DECIMAL(10, 2) NOT NULL,
        category VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    );`

	// Execute the drop query
	if _, err := db.Exec(dropQuery); err != nil {
		log.Fatalf("Could not drop products table: %v", err)
	}
	log.Println("Dropped products table if it existed")

	// Execute the create query
	if _, err := db.Exec(createQuery); err != nil {
		log.Fatalf("Could not create products table: %v", err)
	}
	log.Println("Created products table")
}
