package seeder

import (
	"golang-boilerplate/config"
)

func Seed() {

	db := config.GetDB()

	// Run seeders
	SeedCategoriesTable(db)
	SeedProductsTable(db)
	SeedUsersTable(db)
	SeedMenusTable(db)
	SeedRolesTable(db)      
	SeedPermissionsTable(db)
}
