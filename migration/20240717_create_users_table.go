// migration/20240717_create_users_table.go
package migration

import (
	"database/sql"
	"log"
)

func CreateUsersTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    );`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not create users table: %v", err)
	}
	log.Println("Created users table")
}
