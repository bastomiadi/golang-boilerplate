// seeder/20230717_seed_categories_table.go
package seeder

import (
	"golang-boilerplate/common/models/v1"
	"golang-boilerplate/config"
	"log"
)

func SeedCategoriesTable() {
	db := config.GetDB()

	// Define initial categories
	categories := []models.Category{
		{Name: "Category 1"},
		{Name: "Category 2"},
		{Name: "Category 3"},
	}

	// Perform the seed operation
	if err := db.Create(&categories).Error; err != nil {
		log.Fatalf("Could not seed categories table: %v", err)
	}

	log.Println("Seeded categories table")
}
