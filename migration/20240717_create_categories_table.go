// migration/20240717_create_categories_table.go
package migration

import (
	"database/sql"
	"log"
)

func CreateCategoriesTable(db *sql.DB) {
	dropQuery := `
    DROP TABLE IF EXISTS categories;
    `
	createQuery := `
    CREATE TABLE IF NOT EXISTS categories (
        id INT AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    );`

	// Execute the drop query
	if _, err := db.Exec(dropQuery); err != nil {
		log.Fatalf("Could not drop categories table: %v", err)
	}
	log.Println("Dropped categories table if it existed")

	// Execute the create query
	if _, err := db.Exec(createQuery); err != nil {
		log.Fatalf("Could not create categories table: %v", err)
	}
	log.Println("Created categories table")
}
