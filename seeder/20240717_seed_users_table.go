// seeder/20240717_seed_users_table.go
package seeder

import (
	"database/sql"
	"log"
)

func SeedUsersTable(db *sql.DB) {
	query := `
    INSERT INTO users (name, email, password) VALUES
    ('John Doe', 'john@example.com', 'password123'),
    ('Jane Doe', 'jane@example.com', 'password456');`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Could not seed users table: %v", err)
	}
	log.Println("Seeded users table")
}
