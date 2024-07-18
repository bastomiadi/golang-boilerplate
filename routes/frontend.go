package routes

import (
	frontendControllers "golang-boilerplate/frontend/controllers"
	"net/http"
)

// Add your frontend controllers
func RegisterFrontendRoutes() {
	// Define frontend routes here
	http.HandleFunc("/frontend/", frontendControllers.HandleRequest)
}
