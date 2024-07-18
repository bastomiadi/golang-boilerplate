package routes

import (
	apiV1Controllers "golang-boilerplate/api/v1/controllers"
	apiV2Controllers "golang-boilerplate/api/v2/controllers"
	"net/http"
)

// Add your API controllers

func RegisterAPIRoutes() {
	// Define API routes here
	http.HandleFunc("/api/v1/", apiV1Controllers.HandleRequest)
	http.HandleFunc("/api/v2/", apiV2Controllers.HandleRequest)
}
