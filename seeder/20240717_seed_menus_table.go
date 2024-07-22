// seeder/20240717_seed_menus_table.go
package seeder

import (
	"golang-boilerplate/common/models/v1"
	"golang-boilerplate/config"
	"log"
)

func SeedMenusTable() {
	db := config.GetDB()

	// Define initial menu items
	menus := []models.Menu{
		{Name: "Dashboard", Link: "/backend/dashboard", Parent: 0, Icon: "", DisplayOrder: 0},
		{Name: "Master", Link: "", Parent: 0, Icon: "", DisplayOrder: 0},
		{Name: "Categories", Link: "/backend/categories", Parent: 2, Icon: "", DisplayOrder: 0},
		{Name: "Products", Link: "/backend/products", Parent: 2, Icon: "", DisplayOrder: 0},
		{Name: "RBAC", Link: "", Parent: 0, Icon: "", DisplayOrder: 0},
		{Name: "Users", Link: "/backend/users", Parent: 5, Icon: "", DisplayOrder: 0},
		{Name: "Roles", Link: "/backend/roles", Parent: 5, Icon: "", DisplayOrder: 0},
		{Name: "Permissions", Link: "/backend/permissions", Parent: 5, Icon: "", DisplayOrder: 0},
		{Name: "Menus", Link: "/backend/menus", Parent: 5, Icon: "", DisplayOrder: 0},
	}

	// Perform the seed operation
	if err := db.Create(&menus).Error; err != nil {
		log.Fatalf("Could not seed menus table: %v", err)
	}

	log.Println("Seeded menus table")
}
