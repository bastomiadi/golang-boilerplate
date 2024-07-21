// seeder/20240717_seed_users_table.go
package seeder

import (
	"database/sql"
	"golang-boilerplate/utils"
	"log"
)

func SeedUsersTable(db *sql.DB) {

	password, _ := utils.HashPassword("password")

	query := `
	INSERT INTO users (username, name, email, password) VALUES
	('admin', 'Admin Account', 'admin@admin.com', ?),
	('user', 'User Account', 'user@user.com', ?);`

	if _, err := db.Exec(query, password, password); err != nil {
		log.Fatalf("Could not seed users table: %v", err)
	}
	log.Println("Seeded users table")

	// for i := 0; i < 10; i++ {

	// 	user := models.User{
	// 		Username: faker.Username(),
	// 		Name:     faker.Name(),
	// 		Email:    faker.Email(),
	// 		Password: string(password),
	// 	}

	// 	query := `INSERT INTO users (username, name, email, password) VALUES (?, ?, ?, ?)`
	// 	_, err := db.Exec(query, user.Username, user.Name, user.Email, user.Password)
	// 	if err != nil {
	// 		log.Fatalf("Could not seed users table: %v", err)
	// 	}
	// }

	log.Println("Seeded users table with sample data")
}
