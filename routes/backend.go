package routes

import (
	backendControllers "golang-boilerplate/backend/controllers"
	"net/http"
)

func RegisterBackendRoutes() {

	http.HandleFunc("/backend/register", backendControllers.HandleRegister)
	http.HandleFunc("/backend/login", backendControllers.HandleLogin)
	http.HandleFunc("/backend/dashboard", backendControllers.Dashboard)
	http.HandleFunc("/backend", backendControllers.Dashboard)

	// Initialize routes categories
	http.HandleFunc("/backend/categories", backendControllers.HandleCategoryIndex)
	http.HandleFunc("/backend/categories/list", backendControllers.HandleCategoryList)
	http.HandleFunc("/backend/categories/create", backendControllers.HandleCategoryCreate)
	http.HandleFunc("/backend/categories/update", backendControllers.HandleCategoryUpdate)
	http.HandleFunc("/backend/categories/delete", backendControllers.HandleCategoryDelete)

	// Initializa routes products
	http.HandleFunc("/backend/products", backendControllers.HandleProductIndex)
	http.HandleFunc("/backend/products/list", backendControllers.HandleProductList)
	http.HandleFunc("/backend/products/create", backendControllers.HandleProductCreate)
	http.HandleFunc("/backend/products/update", backendControllers.HandleProductUpdate)
	http.HandleFunc("/backend/products/delete", backendControllers.HandleProductDelete)

	http.HandleFunc("/backend/users", backendControllers.HandleUserIndex)
	http.HandleFunc("/backend/users/list", backendControllers.HandleUserList)
	http.HandleFunc("/backend/users/create", backendControllers.HandleUserCreate)
	http.HandleFunc("/backend/users/update", backendControllers.HandleUserUpdate)
	http.HandleFunc("/backend/users/delete", backendControllers.HandleUserDelete)

	http.HandleFunc("/backend/menus", backendControllers.HandleMenuIndex)
	http.HandleFunc("/backend/menus/list", backendControllers.HandleMenuList)
	http.HandleFunc("/backend/menus/create", backendControllers.HandleMenuCreate)
	http.HandleFunc("/backend/menus/update", backendControllers.HandleMenuUpdate)
	http.HandleFunc("/backend/menus/delete", backendControllers.HandleMenuDelete)

	http.HandleFunc("/backend/roles", backendControllers.HandleRoleIndex)
	http.HandleFunc("/backend/roles/list", backendControllers.HandleRoleList)
	http.HandleFunc("/backend/roles/create", backendControllers.HandleRoleCreate)
	http.HandleFunc("/backend/roles/update", backendControllers.HandleRoleUpdate)
	http.HandleFunc("/backend/roles/delete", backendControllers.HandleRoleDelete)

	http.HandleFunc("/backend/permissions", backendControllers.HandlePermissionIndex)
	http.HandleFunc("/backend/permissions/list", backendControllers.HandlePermissionList)
	http.HandleFunc("/backend/permissions/create", backendControllers.HandlePermissionCreate)
	http.HandleFunc("/backend/permissions/update", backendControllers.HandlePermissionUpdate)
	http.HandleFunc("/backend/permissions/delete", backendControllers.HandlePermissionDelete)
}
