// migration/20240717_create_categories_table.go
package migration

import (
	"database/sql"
	"log"
)

func CreateCategoriesTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS categories (
        id INT AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    );`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not create categories table: %v", err)
	}
	log.Println("Created categories table")
}
