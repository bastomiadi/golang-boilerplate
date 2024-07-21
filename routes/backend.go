package routes

import (
	backendControllers "golang-boilerplate/backend/controllers"
	"golang-boilerplate/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))

func protected(handler http.HandlerFunc) http.Handler {
	return middleware.AuthMiddleware(handler)
}

func RegisterBackendRoutes(router *mux.Router) {
	// Set the session store
	middleware.SetSessionStore(store)

	// Unprotected routes
	router.HandleFunc("/backend/register", backendControllers.HandleRegister)
	router.HandleFunc("/backend/login", backendControllers.HandleLogin)
	router.HandleFunc("/backend/logout", backendControllers.HandleLogout)

	// Protected routes
	router.Handle("/backend/dashboard", protected(backendControllers.Dashboard))
	router.Handle("/backend/", protected(http.HandlerFunc(backendControllers.Dashboard)))

	// Initialize routes categories
	router.Handle("/backend/categories", protected(backendControllers.HandleCategoryIndex))
	router.Handle("/backend/categories/list", protected(backendControllers.HandleCategoryList))
	router.Handle("/backend/categories/create", protected(backendControllers.HandleCategoryCreate))
	router.Handle("/backend/categories/update", protected(backendControllers.HandleCategoryUpdate))
	router.Handle("/backend/categories/delete", protected(backendControllers.HandleCategoryDelete))

	// Initialize routes products
	router.Handle("/backend/products", protected(backendControllers.HandleProductIndex))
	router.Handle("/backend/products/list", protected(backendControllers.HandleProductList))
	router.Handle("/backend/products/create", protected(backendControllers.HandleProductCreate))
	router.Handle("/backend/products/update", protected(backendControllers.HandleProductUpdate))
	router.Handle("/backend/products/delete", protected(backendControllers.HandleProductDelete))

	// Initialize routes users
	router.Handle("/backend/users", protected(backendControllers.HandleUserIndex))
	router.Handle("/backend/users/list", protected(backendControllers.HandleUserList))
	router.Handle("/backend/users/create", protected(backendControllers.HandleUserCreate))
	router.Handle("/backend/users/update", protected(backendControllers.HandleUserUpdate))
	router.Handle("/backend/users/delete", protected(backendControllers.HandleUserDelete))

	// Initialize routes menus
	router.Handle("/backend/menus", protected(backendControllers.HandleMenuIndex))
	router.Handle("/backend/menus/list", protected(backendControllers.HandleMenuList))
	router.Handle("/backend/menus/create", protected(backendControllers.HandleMenuCreate))
	router.Handle("/backend/menus/update", protected(backendControllers.HandleMenuUpdate))
	router.Handle("/backend/menus/delete", protected(backendControllers.HandleMenuDelete))

	// Initialize routes roles
	router.Handle("/backend/roles", protected(backendControllers.HandleRoleIndex))
	router.Handle("/backend/roles/list", protected(backendControllers.HandleRoleList))
	router.Handle("/backend/roles/create", protected(backendControllers.HandleRoleCreate))
	router.Handle("/backend/roles/update", protected(backendControllers.HandleRoleUpdate))
	router.Handle("/backend/roles/delete", protected(backendControllers.HandleRoleDelete))

	// Initialize routes permissions
	router.Handle("/backend/permissions", protected(backendControllers.HandlePermissionIndex))
	router.Handle("/backend/permissions/list", protected(backendControllers.HandlePermissionList))
	router.Handle("/backend/permissions/create", protected(backendControllers.HandlePermissionCreate))
	router.Handle("/backend/permissions/update", protected(backendControllers.HandlePermissionUpdate))
	router.Handle("/backend/permissions/delete", protected(backendControllers.HandlePermissionDelete))
}
