package seeder

func Seed() {
	// Run seeders
	SeedCategoriesTable()
	SeedProductsTable()
	SeedUsersTable()
	SeedMenusTable()
	SeedRolesTable()
	SeedPermissionsTable()
	SeedUserRolesTable()
	SeedRolePermissionsTable()
}
