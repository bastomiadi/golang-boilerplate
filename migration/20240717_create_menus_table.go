// migration/20240717_create_menus_table.go
package migration

import (
	"database/sql"
	"log"
)

func CreateMenusTable(db *sql.DB) {
	dropQuery := `
    DROP TABLE IF EXISTS menus;
    `
	createQuery := `
    CREATE TABLE IF NOT EXISTS menus (
        id INT AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        link VARCHAR(255) NOT NULL,
        parent INT DEFAULT 0,
        icon VARCHAR(50),
        display_order INT DEFAULT 0,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    );`

	// Execute the drop query
	if _, err := db.Exec(dropQuery); err != nil {
		log.Fatalf("Could not drop menus table: %v", err)
	}
	log.Println("Dropped menus table if it existed")

	// Execute the create query
	if _, err := db.Exec(createQuery); err != nil {
		log.Fatalf("Could not create menus table: %v", err)
	}
	log.Println("Created menus table")
}
