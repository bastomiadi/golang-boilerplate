// migration/20240717_create_users_table.go
package migration

import (
	"database/sql"
	"log"
)

func CreateUsersTable(db *sql.DB) {

	// Drop dependent tables first
	dropUserRolesTableQuery := `
    DROP TABLE IF EXISTS user_roles;
    `
	if _, err := db.Exec(dropUserRolesTableQuery); err != nil {
		log.Fatalf("Could not drop user_roles table: %v", err)
	}
	log.Println("Dropped user_roles table if it existed")

	dropQuery := `
    DROP TABLE IF EXISTS users;
    `
	createQuery := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT,
        username VARCHAR(255) NOT NULL,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    );`

	// Execute the drop query
	if _, err := db.Exec(dropQuery); err != nil {
		log.Fatalf("Could not drop users table: %v", err)
	}
	log.Println("Dropped users table if it existed")

	// Execute the create query
	if _, err := db.Exec(createQuery); err != nil {
		log.Fatalf("Could not create users table: %v", err)
	}
	log.Println("Created users table")
}
