package migration

import "golang-boilerplate/config"

func Migrate() {

	db := config.GetDB()

	// Run migrations
	CreateCategoriesTable(db)
	CreateProductsTable(db)
	CreateUsersTable(db)
	CreateMenusTable(db)
	UpdateRolesTable(db)
	UpdatePermissionsTable(db)
	CreateRolePermissionsTable(db)
	CreateUserRolesTable(db)
}
