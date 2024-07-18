// migration/20240717_create_menus_table.go
package migration

import (
	"database/sql"
	"log"
)

func CreateMenusTable(db *sql.DB) {
	query := `
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

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not create menus table: %v", err)
	}
	log.Println("Created menus table")
}
