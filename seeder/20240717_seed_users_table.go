// seeder/20240717_seed_users_table.go
package seeder

import (
	"database/sql"
	"log"
)

func SeedUsersTable(db *sql.DB) {
	query := `
    INSERT INTO users (username, name, email, password) VALUES
    ('admin', 'admin', 'admin@example.com', 'password'),
    ('user1', 'user1', 'user1@example.com', 'password'),
    ('user2', 'user2', 'user2@example.com', 'password');`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not seed users table: %v", err)
	}
	log.Println("Seeded users table")
}
