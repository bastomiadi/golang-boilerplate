package seeder

import (
	"golang-boilerplate/common/models/v1"
	"golang-boilerplate/config"
	"golang-boilerplate/utils"
	"log"
)

func SeedUsersTable() {
	db := config.GetDB()

	// Hash the password
	password, err := utils.HashPassword("password")
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Define initial users
	users := []models.User{
		{Username: "admin", Name: "Admin Account", Email: "admin@admin.com", Password: password},
		{Username: "user", Name: "User Account", Email: "user@user.com", Password: password},
	}

	// Perform the seed operation
	if err := db.Create(&users).Error; err != nil {
		log.Fatalf("Could not seed users table: %v", err)
	}

	// Optional: Seed additional users with sample data
	// for i := 0; i < 10; i++ {
	// 	user := models.User{
	// 		Username: faker.Username(),
	// 		Name:     faker.Name(),
	// 		Email:    faker.Email(),
	// 		Password: password,
	// 	}

	// 	if err := db.Create(&user).Error; err != nil {
	// 		log.Fatalf("Could not seed additional user: %v", err)
	// 	}
	// }

	log.Println("Seeded users table with sample data")
}
