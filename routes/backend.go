package routes

import (
	backendControllers "golang-boilerplate/backend/controllers"
	"golang-boilerplate/middleware"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))

func protected(handler http.HandlerFunc) http.Handler {
	return middleware.AuthMiddleware(handler)
}

func RegisterBackendRoutes() {

	// Set the session store
	middleware.SetSessionStore(store)

	// Unprotected routes
	http.HandleFunc("/backend/register", backendControllers.HandleRegister)
	http.HandleFunc("/backend/login", backendControllers.HandleLogin)
	http.HandleFunc("/backend/logout", backendControllers.HandleLogout)

	// Protected routes
	http.Handle("/backend/dashboard", protected(backendControllers.Dashboard))
	http.Handle("/backend/", protected(http.HandlerFunc(backendControllers.Dashboard)))

	// Initialize routes categories
	http.Handle("/backend/categories", protected(backendControllers.HandleCategoryIndex))
	http.Handle("/backend/categories/list", protected(backendControllers.HandleCategoryList))
	http.Handle("/backend/categories/create", protected(backendControllers.HandleCategoryCreate))
	http.Handle("/backend/categories/update", protected(backendControllers.HandleCategoryUpdate))
	http.Handle("/backend/categories/delete", protected(backendControllers.HandleCategoryDelete))

	// Initialize routes products
	http.Handle("/backend/products", protected(backendControllers.HandleProductIndex))
	http.Handle("/backend/products/list", protected(backendControllers.HandleProductList))
	http.Handle("/backend/products/create", protected(backendControllers.HandleProductCreate))
	http.Handle("/backend/products/update", protected(backendControllers.HandleProductUpdate))
	http.Handle("/backend/products/delete", protected(backendControllers.HandleProductDelete))

	// Initialize routes users
	http.Handle("/backend/users", protected(backendControllers.HandleUserIndex))
	http.Handle("/backend/users/list", protected(backendControllers.HandleUserList))
	http.Handle("/backend/users/create", protected(backendControllers.HandleUserCreate))
	http.Handle("/backend/users/update", protected(backendControllers.HandleUserUpdate))
	http.Handle("/backend/users/delete", protected(backendControllers.HandleUserDelete))

	// Initialize routes menus
	http.Handle("/backend/menus", protected(backendControllers.HandleMenuIndex))
	http.Handle("/backend/menus/list", protected(backendControllers.HandleMenuList))
	http.Handle("/backend/menus/create", protected(backendControllers.HandleMenuCreate))
	http.Handle("/backend/menus/update", protected(backendControllers.HandleMenuUpdate))
	http.Handle("/backend/menus/delete", protected(backendControllers.HandleMenuDelete))

	// Initialize routes roles
	http.Handle("/backend/roles", protected(backendControllers.HandleRoleIndex))
	http.Handle("/backend/roles/list", protected(backendControllers.HandleRoleList))
	http.Handle("/backend/roles/create", protected(backendControllers.HandleRoleCreate))
	http.Handle("/backend/roles/update", protected(backendControllers.HandleRoleUpdate))
	http.Handle("/backend/roles/delete", protected(backendControllers.HandleRoleDelete))

	// Initialize routes permissions
	http.Handle("/backend/permissions", protected(backendControllers.HandlePermissionIndex))
	http.Handle("/backend/permissions/list", protected(backendControllers.HandlePermissionList))
	http.Handle("/backend/permissions/create", protected(backendControllers.HandlePermissionCreate))
	http.Handle("/backend/permissions/update", protected(backendControllers.HandlePermissionUpdate))
	http.Handle("/backend/permissions/delete", protected(backendControllers.HandlePermissionDelete))
}
