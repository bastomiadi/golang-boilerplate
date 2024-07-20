// seeder/20240717_seed_users_table.go
package seeder

import (
	"database/sql"
	"golang-boilerplate/backend/models"
	"golang-boilerplate/utils"
	"log"

	"github.com/go-faker/faker/v4"
)

func SeedUsersTable(db *sql.DB) {
	// query := `
	// INSERT INTO users (username, name, email, password) VALUES
	// ('admin', 'admin', 'admin@example.com', 'password'),
	// ('user1', 'user1', 'user1@example.com', 'password'),
	// ('user2', 'user2', 'user2@example.com', 'password');`

	// if _, err := db.Exec(query); err != nil {
	// 	log.Fatalf("Could not seed users table: %v", err)
	// }
	// log.Println("Seeded users table")

	for i := 0; i < 10; i++ {
		password, _ := utils.HashPassword("password")

		user := models.User{
			Username: faker.Username(),
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: string(password),
		}

		query := `INSERT INTO users (username, name, email, password) VALUES (?, ?, ?, ?)`
		_, err := db.Exec(query, user.Username, user.Name, user.Email, user.Password)
		if err != nil {
			log.Fatalf("Could not seed users table: %v", err)
		}
	}

	log.Println("Seeded users table with sample data")
}
