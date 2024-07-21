package routes

import (
	apiV1 "golang-boilerplate/api/v1/controllers"
	apiV2 "golang-boilerplate/api/v2/controllers"

	"github.com/gorilla/mux"
)

// Add your API controllers
func RegisterApiRoutes(router *mux.Router) {
	v1 := router.PathPrefix("/api/v1").Subrouter()

	// auth
	v1.HandleFunc("/register", apiV1.Register).Methods("POST")
	v1.HandleFunc("/login", apiV1.Login).Methods("POST")
	v1.HandleFunc("/home", apiV1.Home).Methods("GET")

	// categories
	v1.HandleFunc("/categories", apiV1.GetCategories).Methods("GET")
	v1.HandleFunc("/categories/create", apiV1.CreateCategory).Methods("POST")
	v1.HandleFunc("/categories/update", apiV1.UpdateCategory).Methods("PUT")
	v1.HandleFunc("/categories/delete", apiV1.DeleteCategory).Methods("DELETE")

	v2 := router.PathPrefix("/api/v2").Subrouter()
	v2.HandleFunc("/", apiV2.HandleRequest).Methods("GET")
}
