package migration

import "golang-boilerplate/config"

func Migrate() {

	db := config.GetDB()

	// Run migrations
	CreateCategoriesTable(db)
	CreateProductsTable(db)
	CreateUsersTable(db)
	CreateProfilesTable(db)
	CreateMenusTable(db)
	CreateRolesTable(db)
	CreatePermissionsTable(db)
	CreateRolePermissionsTable(db)
	CreateUserRolesTable(db)
}
