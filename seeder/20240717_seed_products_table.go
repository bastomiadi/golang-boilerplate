// seeder/20240717_seed_products_table.go
package seeder

import (
	"golang-boilerplate/common/models/v1"
	"golang-boilerplate/config"
	"log"
)

func SeedProductsTable() {
	db := config.GetDB()

	// Define initial products
	products := []models.Product{
		{Name: "Product 1", Price: 99.99, Category: "Category 1"},
		{Name: "Product 2", Price: 49.99, Category: "Category 2"},
		{Name: "Product 3", Price: 149.99, Category: "Category 1"},
	}

	// Perform the seed operation
	if err := db.Create(&products).Error; err != nil {
		log.Fatalf("Could not seed products table: %v", err)
	}

	log.Println("Seeded products table")
}
